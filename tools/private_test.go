package tools

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	animals "main/Animals"
	"main/abstraction"
	svc "main/service"
	"main/things"
)

/* ============================
   ВСПОМОГАТЕЛЬНЫЕ ТИПЫ/ХЕЛПЕРЫ
   ============================ */

type fakeZoo struct {
	animals   []abstraction.IAlive
	hand      []abstraction.IHerbo
	inventory []abstraction.IInventory
}

func (z *fakeZoo) Admit(a abstraction.IAlive) {
	z.animals = append(z.animals, a)
	if h, ok := a.(abstraction.IHerbo); ok {
		if h.Friendliness() >= svc.MinFriendlinessLevel {
			z.hand = append(z.hand, h)
		}
	}
	if inv, ok := a.(abstraction.IInventory); ok {
		z.inventory = append(z.inventory, inv)
	}
}
func (z *fakeZoo) AddInventory(x abstraction.IInventory) { z.inventory = append(z.inventory, x) }
func (z *fakeZoo) GetAllAnimals() []abstraction.IAlive {
	out := make([]abstraction.IAlive, len(z.animals))
	copy(out, z.animals)
	return out
}
func (z *fakeZoo) GetAllHandAnimal(minFriend int) []abstraction.IHerbo {
	res := make([]abstraction.IHerbo, 0, len(z.hand))
	for _, h := range z.hand {
		if h.Friendliness() >= minFriend {
			res = append(res, h)
		}
	}
	return res
}
func (z *fakeZoo) GetAllInventory() []abstraction.IInventory {
	out := make([]abstraction.IInventory, len(z.inventory))
	copy(out, z.inventory)
	return out
}
func (z *fakeZoo) CountFood() (total int) {
	for _, a := range z.animals {
		total += a.Food()
	}
	return
}
func (z *fakeZoo) NextInventoryNumber() int { return len(z.inventory) + 1 }

type fakeVet struct {
	ok     bool
	reason string
}

func (v *fakeVet) CheckHealth(level svc.HealthLevel) (bool, string) { return v.ok, v.reason }

// перехват stdout на время fn
func captureOutput(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	_ = w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

// конструктор тестовой консоли без чтения меню из файлов
func newTestConsole(z svc.Zoo, v svc.VetClinic) *Console {
	c := &Console{Zoo: z, Vet: v}
	return c
}

/* ============================
   ТЕСТЫ: console.go
   ============================ */

func Test_console_showMenu(t *testing.T) {

	c := newTestConsole(&fakeZoo{}, &fakeVet{ok: true, reason: ""})
	c.items = []MenuItem{{Key: "add", Text: "Добавить"}, {Key: "exit", Text: "Выход"}}
	c.in = bufio.NewScanner(strings.NewReader("\n"))

	out := captureOutput(func() { c.showMenu() })
	if !strings.Contains(out, "Добавить") || !strings.Contains(out, "Выход") {
		t.Fatalf("showMenu не вывел пункты меню, вывод: %q", out)
	}
}

func Test_console_getUserChoice_valid(t *testing.T) {
	c := newTestConsole(&fakeZoo{}, &fakeVet{ok: true, reason: ""})
	c.in = newScanner("2\n")
	got := c.getUserChoice()
	if got != 2 {
		t.Fatalf("ожидали 2, получили %d", got)
	}
}

func Test_console_getUserChoice_invalid(t *testing.T) {
	c := newTestConsole(&fakeZoo{}, &fakeVet{ok: true, reason: ""})
	c.in = newScanner("abc\n")
	got := c.getUserChoice()
	if got != -1 {
		t.Fatalf("ожидали -1 для нечислового ввода, получили %d", got)
	}
}

func Test_console_waitForEnter(t *testing.T) {
	c := newTestConsole(&fakeZoo{}, &fakeVet{ok: true, reason: ""})
	c.in = newScanner("\n") // просто «нажали Enter»
	// не должен зависнуть / паниковать
	c.waitForEnter()
}

func Test_console_changeLanguage_happy(t *testing.T) {
	c := newTestConsole(&fakeZoo{}, &fakeVet{ok: true, reason: ""})
	c.items = []MenuItem{{Key: "exit", Text: "Выход"}} // до смены
	c.in = newScanner("ru\n\n")

	out := captureOutput(func() { c.changeLanguage() })
	if !strings.Contains(out, "Язык успешно изменен") && !strings.Contains(out, "Используется предыдущий") {
		t.Fatalf("ожидали сообщение о смене языка или об ошибке, вывод: %q", out)
	}
}

func Test_console_route_unknown(t *testing.T) {
	c := newTestConsole(&fakeZoo{}, &fakeVet{ok: true, reason: ""})
	c.in = newScanner("\n")
	out := captureOutput(func() { c.route("unknown") })
	if !strings.Contains(out, "Неизвестная функция") {
		t.Fatalf("route должен оповестить об неизвестной функции, вывод: %q", out)
	}
}

func Test_console_Run_with_manual_items_exit(t *testing.T) {
	c := newTestConsole(&fakeZoo{}, &fakeVet{ok: true, reason: ""})
	// вручную задаём меню, чтобы не читать JSON
	c.items = []MenuItem{{Key: "add", Text: "Добавить"}, {Key: "exit", Text: "Выход"}}
	// для выхода Run использует условие choice == len(items)+1
	// len=2 → выбор «3» завершает цикл
	c.in = newScanner("3\n")

	captureOutput(func() { c.Run() }) // не должен паниковать/зависнуть
}

/* ============================
   ТЕСТЫ: flows.go (GetAnimal, stringsTrimSpace)
   ============================ */

func Test_stringsTrimSpace(t *testing.T) {
	in := " \t\n  волк \r\n "
	out := stringsTrimSpace(in)
	if out != "волк" {
		t.Fatalf("ожидали 'волк', получили %q", out)
	}
}

func Test_GetAnimal_Monkey_OK(t *testing.T) {
	z := &fakeZoo{}
	v := &fakeVet{ok: true}
	c := newTestConsole(z, v)

	// тип, здоровье, еда, дружелюбность, имя
	c.in = newScanner("Обезьяна\n5\n10\n7\nБобо\n")

	captureOutput(func() { c.GetAnimal() })

	if len(z.animals) != 1 {
		t.Fatalf("ожидали 1 животное, получили %d", len(z.animals))
	}
	if len(z.inventory) != 1 {
		t.Fatalf("животное должно быть добавлено и как инвентарь, сейчас %d", len(z.inventory))
	}
	// дружелюбность 7 → попадёт в контактный
	if len(z.hand) != 1 {
		t.Fatalf("ожидали 1 животное в контактном зоопарке, получили %d", len(z.hand))
	}
}

func Test_GetAnimal_Tiger_OK(t *testing.T) {
	z := &fakeZoo{}
	v := &fakeVet{ok: true}
	c := newTestConsole(z, v)

	c.in = newScanner("Тигр\n4\n20\nБарс\n")

	captureOutput(func() { c.GetAnimal() })

	if len(z.animals) != 1 {
		t.Fatalf("ожидали 1 животное, получили %d", len(z.animals))
	}
	// хищник не должен попасть в контактный (он не IHerbo)
	if len(z.hand) != 0 {
		t.Fatalf("тигр не должен быть в контактном зоопарке")
	}
}

func Test_GetAnimal_ClinicReject(t *testing.T) {
	z := &fakeZoo{}
	v := &fakeVet{ok: false, reason: "низкое здоровье"}
	c := newTestConsole(z, v)

	c.in = newScanner("Кролик\n1\n") // клиника отклонит по здоровью

	out := captureOutput(func() { c.GetAnimal() })
	if !strings.Contains(out, "Отказ клиники") {
		t.Fatalf("ожидали отказ клиники, вывод: %q", out)
	}
	if len(z.animals) != 0 {
		t.Fatalf("животное не должно быть принято при отказе клиники")
	}
}

/* ============================
   ТЕСТЫ: input.go
   ============================ */

func Test_GetName(t *testing.T) {
	c := newTestConsole(&fakeZoo{}, &fakeVet{ok: true})
	c.in = newScanner("Мурзик\n")
	name, err := c.GetName()
	if err != nil || name != "Мурзик" {
		t.Fatalf("GetName ошибка: name=%q err=%v", name, err)
	}
}

func Test_GetValidFood(t *testing.T) {
	c := newTestConsole(&fakeZoo{}, &fakeVet{ok: true})
	c.in = newScanner("12\n")
	v, err := c.GetValidFood()
	if err != nil || v != 12 {
		t.Fatalf("ожидали 12, получили %d (err=%v)", v, err)
	}
}

func Test_GetFriendless(t *testing.T) {
	c := newTestConsole(&fakeZoo{}, &fakeVet{ok: true})
	c.in = newScanner("9\n")
	v, err := c.GetFriendless()
	if err != nil || v != 9 {
		t.Fatalf("ожидали 9, получили %d (err=%v)", v, err)
	}
}

func Test_GetHealthLevel(t *testing.T) {
	c := newTestConsole(&fakeZoo{}, &fakeVet{ok: true})
	c.in = newScanner("5\n")
	lv, err := c.GetHealthLevel()
	if err != nil || lv != svc.Health5 {
		t.Fatalf("ожидали Health5, получили %v (err=%v)", lv, err)
	}
}

func Test_handleErr(t *testing.T) {
	c := newTestConsole(&fakeZoo{}, &fakeVet{ok: true})
	c.in = bufio.NewScanner(strings.NewReader("\n"))

	out := captureOutput(func() {
		if ok := c.handleErr(io.EOF); !ok {
			t.Fatal("handleErr(true) должен вернуть true")
		}
	})
	if !strings.Contains(out, "EOF") {
		t.Fatalf("ожидали вывод ошибки, получили: %q", out)
	}
	if ok := c.handleErr(nil); ok {
		t.Fatal("handleErr(nil) должен вернуть false")
	}
}

/* ============================
   ТЕСТЫ: things.go (GetThing)
   ============================ */

func Test_GetThing_Computer(t *testing.T) {
	z := &fakeZoo{}
	v := &fakeVet{ok: true}
	c := newTestConsole(z, v)

	// тип, модель, RAM, CPU
	c.in = newScanner("Computer\nHP\n8\ni5\n")

	captureOutput(func() { c.GetThing() })

	if len(z.inventory) != 1 {
		t.Fatalf("ожидали 1 предмет в инвентаре, получили %d", len(z.inventory))
	}
	if z.inventory[0].Number() != 1 {
		t.Fatalf("ожидали номер 1 для первого предмета")
	}
}

func Test_GetThing_Table(t *testing.T) {
	z := &fakeZoo{}
	v := &fakeVet{ok: true}
	c := newTestConsole(z, v)

	// тип, материал, цвет
	c.in = newScanner("Table\nДерево\nКоричневый\n")
	captureOutput(func() { c.GetThing() })

	if len(z.inventory) != 1 {
		t.Fatalf("ожидали 1 предмет, получили %d", len(z.inventory))
	}
}

func Test_GetThing_Thing(t *testing.T) {
	z := &fakeZoo{}
	v := &fakeVet{ok: true}
	c := newTestConsole(z, v)

	// тип, имя, вес
	c.in = newScanner("Thing\nКоробка\n2.5\n")
	captureOutput(func() { c.GetThing() })

	if len(z.inventory) != 1 {
		t.Fatalf("ожидали 1 предмет, получили %d", len(z.inventory))
	}
}

/* ============================
   ТЕСТЫ: render.go
   ============================ */

func Test_WriteAllAnimal_and_Hand_and_Inventory(t *testing.T) {
	z := &fakeZoo{}
	v := &fakeVet{ok: true}
	c := newTestConsole(z, v)
	c.in = bufio.NewScanner(strings.NewReader("\n"))

	// добавим несколько сущностей в зоопарк/инвентарь
	z.Admit(animals.NewMonkey("Чита", 10, 8, 1)) // hand
	z.Admit(animals.NewRabbit("Банни", 5, 9, 2)) // hand
	z.Admit(animals.NewTiger("Раджа", 25, 3))    // не hand
	z.AddInventory(things.NewComputer(4, "Lenovo", 16, "i5"))

	out1 := captureOutput(func() { c.WriteAllAnimal() })
	if !strings.Contains(out1, "Общее количество животных: 3") {
		t.Fatalf("ожидали 3 животных, вывод: %q", out1)
	}

	out2 := captureOutput(func() { c.WriteAllHandAnimal() })
	if !strings.Contains(out2, "Общее количество животных в контактном зоопарке: 2") {
		t.Fatalf("ожидали 2 hand-животных, вывод: %q", out2)
	}

	out3 := captureOutput(func() { c.WriteAllInventory() })
	if !strings.Contains(out3, "Общий баланс зоопарка: 4") { // 3 животных + 1 предмет
		t.Fatalf("ожидали 4 предмета инвентаря, вывод: %q", out3)
	}
}

/* ============================
   ХЕЛПЕР: сканер по строке
   ============================ */

func newScanner(input string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(input))
}

// bytes.ReaderScanner — мини-адаптер, чтобы не подключать bufio.Scanner напрямую в тестах.
// Реализация совместима со scanner.Scan()/Text() в консоли.

type ReaderScanner interface {
	Scan() bool
	Text() string
}

type readerScanner struct {
	r *strings.Reader
	b strings.Builder
}

func (rs *readerScanner) Scan() bool {
	rs.b.Reset()
	for {
		ch, _, err := rs.r.ReadRune()
		if err != nil {
			return rs.b.Len() > 0
		}
		if ch == '\n' {
			return true
		}
		rs.b.WriteRune(ch)
	}
}

func (rs *readerScanner) Text() string { return rs.b.String() }

// упрощённый конструктор
func NewReaderScannerString(s string) *readerScanner { return &readerScanner{r: strings.NewReader(s)} }

// совместимость имени
func NewReaderScanner(b []byte) *readerScanner {
	return &readerScanner{r: strings.NewReader(string(b))}
}

// алиасы под краткое создание
func bytesNewReaderScanner(b []byte) *readerScanner       { return NewReaderScanner(b) }
func bytesNewReaderScannerString(s string) *readerScanner { return NewReaderScanner([]byte(s)) }

// нужен для newScanner
type bytesReaderScanner = readerScanner

func bytesNewReaderScannerCompat(b []byte) *bytesReaderScanner { return NewReaderScanner(b) }

func bytesNewReaderScannerStringCompat(s string) *bytesReaderScanner {
	return NewReaderScanner([]byte(s))
}

// factory для newScanner
func bytes_NewReaderScanner(b []byte) *bytesReaderScanner { return NewReaderScanner(b) }

// конечная обёртка, которую мы выше дергаем
func NewBytesReaderScanner(b []byte) *bytesReaderScanner { return NewReaderScanner(b) }

// короткая функция, используемая в тестах выше
func NewScannerString(s string) *bytesReaderScanner { return NewReaderScanner([]byte(s)) }
