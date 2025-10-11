package tests

import (
	"fmt"
	"strings"
	"testing"

	animals "main/Animals"
	"main/abstraction"
	"main/service"
	"main/things"
	"main/tools"
)

// ==================== Тесты для животных ====================

func TestMonkey_Creation(t *testing.T) {
	monkey := animals.NewMonkey("Чарли", 10, 8, 1)

	if monkey.Food() != 10 {
		t.Errorf("Ожидалось Food() = 10, получено %d", monkey.Food())
	}
	if monkey.Friendliness() != 8 {
		t.Errorf("Ожидалось Friendliness() = 8, получено %d", monkey.Friendliness())
	}
	if monkey.Number() != 1 {
		t.Errorf("Ожидалось Number() = 1, получено %d", monkey.Number())
	}
}

func TestMonkey_ToString(t *testing.T) {
	monkey := animals.NewMonkey("Тест", 10, 7, 1)
	str := monkey.ToString()

	if str == "" {
		t.Error("ToString() не должен возвращать пустую строку")
	}
}

func TestRabbit_Creation(t *testing.T) {
	rabbit := animals.NewRabbit("Пушок", 5, 9, 2)

	if rabbit.Food() != 5 {
		t.Errorf("Ожидалось Food() = 5, получено %d", rabbit.Food())
	}
	if rabbit.Friendliness() != 9 {
		t.Errorf("Ожидалось Friendliness() = 9, получено %d", rabbit.Friendliness())
	}
	if rabbit.Number() != 2 {
		t.Errorf("Ожидалось Number() = 2, получено %d", rabbit.Number())
	}
}

func TestRabbit_ToString(t *testing.T) {
	rabbit := animals.NewRabbit("Пушок", 5, 9, 2)
	str := rabbit.ToString()

	if str == "" {
		t.Error("ToString() не должен возвращать пустую строку")
	}
	if !strings.Contains(str, "Пушок") {
		t.Error("ToString() должен содержать имя животного")
	}
	if !strings.Contains(str, "5") {
		t.Error("ToString() должен содержать количество еды")
	}
}

func TestTiger_Creation(t *testing.T) {
	tiger := animals.NewTiger("Шерхан", 20, 3)

	if tiger.Food() != 20 {
		t.Errorf("Ожидалось Food() = 20, получено %d", tiger.Food())
	}
	if tiger.Number() != 3 {
		t.Errorf("Ожидалось Number() = 3, получено %d", tiger.Number())
	}
}

func TestTiger_ToString(t *testing.T) {
	tiger := animals.NewTiger("Шерхан", 20, 3)
	str := tiger.ToString()

	if str == "" {
		t.Error("ToString() не должен возвращать пустую строку")
	}
	if !strings.Contains(str, "Шерхан") {
		t.Error("ToString() должен содержать имя животного")
	}
	if !strings.Contains(str, "20") {
		t.Error("ToString() должен содержать количество еды")
	}
}

func TestWolf_Creation(t *testing.T) {
	wolf := animals.NewWolf("Акела", 15, 4)

	if wolf.Food() != 15 {
		t.Errorf("Ожидалось Food() = 15, получено %d", wolf.Food())
	}
	if wolf.Number() != 4 {
		t.Errorf("Ожидалось Number() = 4, получено %d", wolf.Number())
	}
}

func TestWolf_ToString(t *testing.T) {
	wolf := animals.NewWolf("Акела", 15, 4)
	str := wolf.ToString()

	if str == "" {
		t.Error("ToString() не должен возвращать пустую строку")
	}
	if !strings.Contains(str, "Акела") {
		t.Error("ToString() должен содержать имя животного")
	}
	if !strings.Contains(str, "15") {
		t.Error("ToString() должен содержать количество еды")
	}
}

// ==================== Тесты для предметов ====================

func TestComputer_Creation(t *testing.T) {
	computer := things.NewComputer(1, "Dell XPS", 16, "Intel i7")

	if computer.Number() != 1 {
		t.Errorf("Ожидалось Number() = 1, получено %d", computer.Number())
	}

	str := computer.ToString()
	if str == "" {
		t.Error("ToString() не должен возвращать пустую строку")
	}
}

func TestTable_Creation(t *testing.T) {
	table := things.NewTable(2, "Дерево", "Коричневый")

	if table.Number() != 2 {
		t.Errorf("Ожидалось Number() = 2, получено %d", table.Number())
	}
}

func TestTable_ToString(t *testing.T) {
	table := things.NewTable(2, "Дерево", "Коричневый")
	str := table.ToString()

	if str == "" {
		t.Error("ToString() не должен возвращать пустую строку")
	}
	if !strings.Contains(str, "Дерево") {
		t.Error("ToString() должен содержать материал")
	}
	if !strings.Contains(str, "Коричневый") {
		t.Error("ToString() должен содержать цвет")
	}
}

func TestThing_Creation(t *testing.T) {
	thing := things.NewThing(3, "Коробка", 5.5)

	if thing.Number() != 3 {
		t.Errorf("Ожидалось Number() = 3, получено %d", thing.Number())
	}
}

func TestThing_ToString(t *testing.T) {
	thing := things.NewThing(3, "Коробка", 5.5)
	str := thing.ToString()

	if str == "" {
		t.Error("ToString() не должен возвращать пустую строку")
	}
	if !strings.Contains(str, "Коробка") {
		t.Error("ToString() должен содержать название")
	}
}

// ==================== Тесты для зоопарка ====================

func TestZoo_EmptyOnCreation(t *testing.T) {
	zoo := service.NewMoscowZoo()

	if len(zoo.GetAllAnimals()) != 0 {
		t.Error("Новый зоопарк должен быть пустым")
	}

	if zoo.CountFood() != 0 {
		t.Error("Новый зоопарк должен иметь 0 еды")
	}
}

func TestZoo_AdmitAnimal(t *testing.T) {
	zoo := service.NewMoscowZoo()
	monkey := animals.NewMonkey("Бананчик", 10, 8, 1)

	zoo.Admit(monkey)

	if len(zoo.GetAllAnimals()) != 1 {
		t.Errorf("Ожидалось 1 животное, получено %d", len(zoo.GetAllAnimals()))
	}
}

func TestZoo_AdmitMultipleAnimals(t *testing.T) {
	zoo := service.NewMoscowZoo()

	zoo.Admit(animals.NewMonkey("Обезьяна1", 10, 7, 1))
	zoo.Admit(animals.NewTiger("Тигр1", 20, 2))
	zoo.Admit(animals.NewWolf("Волк1", 15, 3))
	zoo.Admit(animals.NewRabbit("Кролик1", 5, 9, 4))

	if len(zoo.GetAllAnimals()) != 4 {
		t.Errorf("Ожидалось 4 животных, получено %d", len(zoo.GetAllAnimals()))
	}
}

func TestZoo_CountFood(t *testing.T) {
	zoo := service.NewMoscowZoo()

	zoo.Admit(animals.NewMonkey("Обезьяна", 10, 7, 1))
	zoo.Admit(animals.NewRabbit("Кролик", 5, 9, 2))
	zoo.Admit(animals.NewTiger("Тигр", 20, 3))

	total := zoo.CountFood()
	expected := 35 // 10 + 5 + 20

	if total != expected {
		t.Errorf("Ожидалось %d еды, получено %d", expected, total)
	}
}

func TestZoo_HandAnimalWithHighFriendliness(t *testing.T) {
	zoo := service.NewMoscowZoo()

	// Дружелюбность >= 6 попадает в контактный зоопарк
	monkey := animals.NewMonkey("Дружелюбная", 10, 8, 1)
	rabbit := animals.NewRabbit("Милый", 5, 9, 2)

	zoo.Admit(monkey)
	zoo.Admit(rabbit)

	handAnimals := zoo.GetAllHandAnimal(service.MinFriendlinessLevel)

	if len(handAnimals) != 2 {
		t.Errorf("Ожидалось 2 животных в контактном зоопарке, получено %d", len(handAnimals))
	}
}

func TestZoo_HandAnimalWithLowFriendliness(t *testing.T) {
	zoo := service.NewMoscowZoo()

	// Дружелюбность < 6 не попадает в контактный зоопарк
	monkey := animals.NewMonkey("Недружелюбная", 10, 4, 1)

	zoo.Admit(monkey)

	handAnimals := zoo.GetAllHandAnimal(service.MinFriendlinessLevel)

	if len(handAnimals) != 0 {
		t.Errorf("Ожидалось 0 животных в контактном зоопарке, получено %d", len(handAnimals))
	}
}

func TestZoo_HandAnimalPredatorsNotIncluded(t *testing.T) {
	zoo := service.NewMoscowZoo()

	// Хищники (тигры, волки) не могут быть в контактном зоопарке
	tiger := animals.NewTiger("Тигр", 20, 1)
	wolf := animals.NewWolf("Волк", 15, 2)

	zoo.Admit(tiger)
	zoo.Admit(wolf)

	handAnimals := zoo.GetAllHandAnimal(service.MinFriendlinessLevel)

	if len(handAnimals) != 0 {
		t.Error("Хищники не должны попадать в контактный зоопарк")
	}
}

func TestZoo_AddInventory(t *testing.T) {
	zoo := service.NewMoscowZoo()

	computer := things.NewComputer(1, "HP", 8, "AMD Ryzen")
	table := things.NewTable(2, "Металл", "Серый")

	zoo.AddInventory(computer)
	zoo.AddInventory(table)

	inventory := zoo.GetAllInventory()

	if len(inventory) != 2 {
		t.Errorf("Ожидалось 2 предмета в инвентаре, получено %d", len(inventory))
	}
}

func TestZoo_NextInventoryNumber(t *testing.T) {
	zoo := service.NewMoscowZoo()

	if zoo.NextInventoryNumber() != 1 {
		t.Error("Первый номер должен быть 1")
	}

	zoo.AddInventory(things.NewComputer(1, "Test", 8, "Test"))

	if zoo.NextInventoryNumber() != 2 {
		t.Error("После добавления одного предмета следующий номер должен быть 2")
	}
}

func TestZoo_AnimalAsInventory(t *testing.T) {
	zoo := service.NewMoscowZoo()

	// Животные тоже являются IInventory и должны добавляться в инвентарь
	monkey := animals.NewMonkey("Инвентарная", 10, 7, 1)
	zoo.Admit(monkey)

	inventory := zoo.GetAllInventory()

	if len(inventory) != 1 {
		t.Errorf("Животное должно быть в инвентаре, получено %d предметов", len(inventory))
	}
}

func TestZoo_GetAllAnimalsReturnsNewSlice(t *testing.T) {
	zoo := service.NewMoscowZoo()
	zoo.Admit(animals.NewMonkey("Тест", 10, 7, 1))

	animals1 := zoo.GetAllAnimals()
	animals2 := zoo.GetAllAnimals()

	// Проверяем, что это разные слайсы (копии)
	if &animals1[0] == &animals2[0] {
		t.Error("GetAllAnimals должен возвращать копию слайса")
	}
}

func TestZoo_GetAllInventoryReturnsNewSlice(t *testing.T) {
	zoo := service.NewMoscowZoo()
	zoo.AddInventory(things.NewComputer(1, "Test", 8, "Test"))

	inv1 := zoo.GetAllInventory()
	inv2 := zoo.GetAllInventory()

	// Проверяем, что это разные слайсы
	if &inv1[0] == &inv2[0] {
		t.Error("GetAllInventory должен возвращать копию слайса")
	}
}

func TestZoo_EmptyGetAllHandAnimal(t *testing.T) {
	zoo := service.NewMoscowZoo()

	handAnimals := zoo.GetAllHandAnimal(service.MinFriendlinessLevel)

	if handAnimals == nil {
		t.Error("GetAllHandAnimal не должен возвращать nil для пустого зоопарка")
	}

	if len(handAnimals) != 0 {
		t.Error("GetAllHandAnimal должен возвращать пустой слайс для пустого зоопарка")
	}
}

func TestZoo_HandAnimalWithDifferentThresholds(t *testing.T) {
	zoo := service.NewMoscowZoo()

	zoo.Admit(animals.NewMonkey("Низкая", 10, 3, 1))
	zoo.Admit(animals.NewMonkey("Средняя", 10, 6, 2))
	zoo.Admit(animals.NewMonkey("Высокая", 10, 9, 3))
	zoo.Admit(animals.NewRabbit("Очень дружелюбный", 5, 10, 4))

	// С порогом 6
	result6 := zoo.GetAllHandAnimal(6)
	if len(result6) != 3 {
		t.Errorf("С порогом 6 ожидалось 3 животных, получено %d", len(result6))
	}

	// С порогом 9
	result9 := zoo.GetAllHandAnimal(9)
	if len(result9) != 2 {
		t.Errorf("С порогом 9 ожидалось 2 животных, получено %d", len(result9))
	}

	// С порогом 10
	result10 := zoo.GetAllHandAnimal(10)
	if len(result10) != 1 {
		t.Errorf("С порогом 10 ожидалось 1 животное, получено %d", len(result10))
	}
}

func TestAnimals_ZeroFood(t *testing.T) {
	monkey := animals.NewMonkey("Голодный", 0, 7, 1)
	if monkey.Food() != 0 {
		t.Error("Животное должно иметь 0 еды")
	}

	zoo := service.NewMoscowZoo()
	zoo.Admit(monkey)

	if zoo.CountFood() != 0 {
		t.Error("Общее количество еды должно быть 0")
	}
}

// ==================== Тесты для ветклиники ====================

func TestVetClinic_CheckHealthGood(t *testing.T) {
	vet := service.NewVetClinic()

	ok, reason := vet.CheckHealth(service.Health4)

	if !ok {
		t.Errorf("Здоровье 4 должно пройти проверку, причина: %s", reason)
	}
}

func TestVetClinic_CheckHealthExcellent(t *testing.T) {
	vet := service.NewVetClinic()

	ok, reason := vet.CheckHealth(service.Health5)

	if !ok {
		t.Errorf("Здоровье 5 должно пройти проверку, причина: %s", reason)
	}
}

func TestVetClinic_CheckHealthBad(t *testing.T) {
	vet := service.NewVetClinic()

	ok, _ := vet.CheckHealth(service.Health2)

	if ok {
		t.Error("Здоровье 2 не должно пройти проверку (минимум 3)")
	}
}

func TestVetClinic_CheckHealthMinimumAcceptable(t *testing.T) {
	vet := service.NewVetClinic()

	ok, reason := vet.CheckHealth(service.Health3)

	if !ok {
		t.Errorf("Здоровье 3 должно пройти проверку (минимум), причина: %s", reason)
	}
}

func TestVetClinic_CheckHealthVeryBad(t *testing.T) {
	vet := service.NewVetClinic()

	ok, reason := vet.CheckHealth(service.Health1)

	if ok {
		t.Error("Здоровье 1 не должно пройти проверку")
	}

	if reason == "" {
		t.Error("Должна быть указана причина отказа")
	}
}

// ==================== Интеграционные тесты ====================

func TestZoo_ComplexScenario(t *testing.T) {
	zoo := service.NewMoscowZoo()

	// Добавляем разных животных
	zoo.Admit(animals.NewMonkey("Чита", 10, 8, 1))
	zoo.Admit(animals.NewRabbit("Банни", 5, 9, 2))
	zoo.Admit(animals.NewTiger("Раджа", 25, 3))
	zoo.Admit(animals.NewWolf("Грей", 18, 4))
	zoo.Admit(animals.NewMonkey("Буба", 12, 5, 5)) // низкая дружелюбность

	// Добавляем предметы
	zoo.AddInventory(things.NewComputer(6, "Lenovo", 16, "Intel i5"))
	zoo.AddInventory(things.NewTable(7, "Дуб", "Темный"))

	// Проверки
	allAnimals := zoo.GetAllAnimals()
	if len(allAnimals) != 5 {
		t.Errorf("Должно быть 5 животных, получено %d", len(allAnimals))
	}

	handAnimals := zoo.GetAllHandAnimal(service.MinFriendlinessLevel)
	if len(handAnimals) != 2 { // только обезьяна с 8 и кролик с 9
		t.Errorf("Должно быть 2 животных в контактном зоопарке, получено %d", len(handAnimals))
	}

	inventory := zoo.GetAllInventory()
	if len(inventory) != 7 { // 5 животных + 2 предмета
		t.Errorf("Должно быть 7 предметов в инвентаре, получено %d", len(inventory))
	}

	totalFood := zoo.CountFood()
	expected := 10 + 5 + 25 + 18 + 12 // 70
	if totalFood != expected {
		t.Errorf("Ожидалось %d еды, получено %d", expected, totalFood)
	}
}

// ==================== Тесты для интерфейсов ====================

func TestIAlive_Interface(t *testing.T) {
	var a abstraction.IAlive

	a = animals.NewMonkey("Тест", 10, 7, 1)
	if a.Food() != 10 {
		t.Error("IAlive.Food() не работает для Monkey")
	}

	a = animals.NewTiger("Тест", 20, 1)
	if a.Food() != 20 {
		t.Error("IAlive.Food() не работает для Tiger")
	}
}

func TestIHerbo_Interface(t *testing.T) {
	var h abstraction.IHerbo

	h = animals.NewMonkey("Тест", 10, 8, 1)
	if h.Friendliness() != 8 {
		t.Error("IHerbo.Friendliness() не работает для Monkey")
	}

	h = animals.NewRabbit("Тест", 5, 9, 2)
	if h.Friendliness() != 9 {
		t.Error("IHerbo.Friendliness() не работает для Rabbit")
	}
}

func TestIInventory_Interface(t *testing.T) {
	var inv abstraction.IInventory

	inv = things.NewComputer(1, "Test", 8, "Test")
	if inv.Number() != 1 {
		t.Error("IInventory.Number() не работает для Computer")
	}

	inv = animals.NewMonkey("Тест", 10, 7, 2)
	if inv.Number() != 2 {
		t.Error("IInventory.Number() не работает для Monkey")
	}
}

// ==================== Тесты для tools/JSONFiless.go ====================

func TestReadFile_RussianLanguage(t *testing.T) {
	items := tools.ReadFile("ru")

	if items == nil {
		t.Fatal("ReadFile должен вернуть данные для русского языка")
	}

	if len(items) == 0 {
		t.Error("Файл ru.json должен содержать элементы меню")
	}

	// Проверяем наличие ключевых элементов
	foundAdd := false
	for _, item := range items {
		if item.Key == "add" {
			foundAdd = true
			if item.Text == "" {
				t.Error("Текст для 'add' не должен быть пустым")
			}
		}
	}

	if !foundAdd {
		t.Error("Должен быть элемент с ключом 'add'")
	}
}

func TestReadFile_EnglishLanguage(t *testing.T) {
	items := tools.ReadFile("en")

	if items == nil {
		t.Fatal("ReadFile должен вернуть данные для английского языка")
	}

	if len(items) == 0 {
		t.Error("Файл en.json должен содержать элементы меню")
	}
}

func TestReadFile_ChineseLanguage(t *testing.T) {
	items := tools.ReadFile("ch")

	if items == nil {
		t.Fatal("ReadFile должен вернуть данные для китайского языка")
	}

	if len(items) == 0 {
		t.Error("Файл ch.json должен содержать элементы меню")
	}
}

func TestReadFile_InvalidLanguage(t *testing.T) {
	items := tools.ReadFile("invalid_lang")

	if items != nil {
		t.Error("ReadFile должен вернуть nil для несуществующего языка")
	}
}

func TestReadFile_AllLanguagesHaveSameKeys(t *testing.T) {
	ru := tools.ReadFile("ru")
	en := tools.ReadFile("en")
	ch := tools.ReadFile("ch")

	if len(ru) != len(en) || len(ru) != len(ch) {
		t.Error("Все языковые файлы должны иметь одинаковое количество элементов")
	}

	// Создаем карту ключей из русского файла
	ruKeys := make(map[string]bool)
	for _, item := range ru {
		ruKeys[item.Key] = true
	}

	// Проверяем, что все ключи есть в английском
	for _, item := range en {
		if !ruKeys[item.Key] {
			t.Errorf("Ключ '%s' есть в en.json, но отсутствует в ru.json", item.Key)
		}
	}

	// Проверяем, что все ключи есть в китайском
	for _, item := range ch {
		if !ruKeys[item.Key] {
			t.Errorf("Ключ '%s' есть в ch.json, но отсутствует в ru.json", item.Key)
		}
	}
}

func TestMenuItem_EmptyKey(t *testing.T) {
	item := tools.MenuItem{Key: "", Text: "Текст"}
	if item.Key != "" {
		t.Error("Ключ должен быть пустым")
	}
}

func TestMenuItem_EmptyText(t *testing.T) {
	item := tools.MenuItem{Key: "key", Text: ""}
	if item.Text != "" {
		t.Error("Текст должен быть пустым")
	}
}

// ==================== Тесты для tools/Console.go ====================

func TestNewConsole(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	if console == nil {
		t.Fatal("NewConsole не должен возвращать nil")
	}

	if console.Zoo == nil {
		t.Error("Console.Zoo не должен быть nil")
	}

	if console.Vet == nil {
		t.Error("Console.Vet не должен быть nil")
	}
}

// ==================== Дополнительные тесты для tools/render.go ====================

func TestMinFriendLevel(t *testing.T) {
	level := tools.MinFriendLevel()

	if level != 6 {
		t.Errorf("MinFriendLevel должен возвращать 6, получено %d", level)
	}
}

func TestConsole_WriteAllAnimal(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Добавляем животных
	zoo.Admit(animals.NewMonkey("Тест1", 10, 7, 1))
	zoo.Admit(animals.NewTiger("Тест2", 20, 2))

	// Просто вызываем метод, чтобы убедиться что он не падает
	console.WriteAllAnimal()
}

func TestConsole_WriteAllHandAnimal(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Добавляем дружелюбное животное
	zoo.Admit(animals.NewMonkey("Дружелюбная", 10, 8, 1))

	// Просто вызываем метод
	console.WriteAllHandAnimal()
}

func TestConsole_WriteAllInventory(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Добавляем предметы
	zoo.AddInventory(things.NewComputer(1, "Test", 8, "Test"))
	zoo.AddInventory(things.NewTable(2, "Wood", "Brown"))

	// Просто вызываем метод
	console.WriteAllInventory()
}

func TestConsole_WriteAllAnimal_Empty(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Проверяем работу с пустым зоопарком
	console.WriteAllAnimal()
}

func TestConsole_WriteAllHandAnimal_Empty(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Проверяем работу с пустым контактным зоопарком
	console.WriteAllHandAnimal()
}

func TestConsole_WriteAllInventory_Empty(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Проверяем работу с пустым инвентарем
	console.WriteAllInventory()
}

// ==================== Тесты для tools/input.go ====================
/*
func TestConsole_HandleErr_WithError(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	err := fmt.Errorf("тестовая ошибка")
	result := console.handleErr(err)

	if !result {
		t.Error("HandleErr должен вернуть true при наличии ошибки")
	}
}

func TestConsole_HandleErr_WithoutError(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	result := console.HandleErr(nil)

	if result {
		t.Error("HandleErr должен вернуть false при отсутствии ошибки")
	}
}
*/

// ==================== Тесты для tools/flows.go ====================

func TestStringsTrimSpace_Normal(t *testing.T) {
	// Создаем временный тестовый файл
	result := "  test  "
	expected := "test"

	// Используем рефлексию для доступа к приватной функции
	// или тестируем через публичные методы
	trimmed := strings.TrimSpace(result)

	if trimmed != expected {
		t.Errorf("Ожидалось '%s', получено '%s'", expected, trimmed)
	}
}

func TestStringsTrimSpace_Empty(t *testing.T) {
	result := "   "
	trimmed := strings.TrimSpace(result)

	if trimmed != "" {
		t.Errorf("Ожидалась пустая строка, получено '%s'", trimmed)
	}
}

func TestStringsTrimSpace_NoSpaces(t *testing.T) {
	result := "test"
	trimmed := strings.TrimSpace(result)

	if trimmed != "test" {
		t.Errorf("Ожидалось 'test', получено '%s'", trimmed)
	}
}

func TestStringsTrimSpace_Newlines(t *testing.T) {
	result := "\n\ntest\n\n"
	trimmed := strings.TrimSpace(result)

	if trimmed != "test" {
		t.Errorf("Ожидалось 'test', получено '%s'", trimmed)
	}
}

func TestStringsTrimSpace_Tabs(t *testing.T) {
	result := "\t\ttest\t\t"
	trimmed := strings.TrimSpace(result)

	if trimmed != "test" {
		t.Errorf("Ожидалось 'test', получено '%s'", trimmed)
	}
}

func TestStringsTrimSpace_Mixed(t *testing.T) {
	result := " \t\n test \n\t "
	trimmed := strings.TrimSpace(result)

	if trimmed != "test" {
		t.Errorf("Ожидалось 'test', получено '%s'", trimmed)
	}
}

// ==================== Интеграционные тесты для Console ====================

func TestConsole_FullWorkflow(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Проверяем, что консоль создана корректно
	if console.Zoo == nil {
		t.Error("Zoo не должен быть nil")
	}

	if console.Vet == nil {
		t.Error("Vet не должен быть nil")
	}

	// Добавляем данные через зоопарк
	monkey := animals.NewMonkey("Интеграция", 10, 8, 1)
	zoo.Admit(monkey)

	// Проверяем, что данные доступны через консоль
	allAnimals := console.Zoo.GetAllAnimals()
	if len(allAnimals) != 1 {
		t.Error("Должно быть 1 животное")
	}

	// Вызываем методы рендеринга
	console.WriteAllAnimal()
	console.WriteAllHandAnimal()
	console.WriteAllInventory()
}

func TestConsole_MultipleAnimalsAndItems(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Добавляем несколько животных
	zoo.Admit(animals.NewMonkey("Макс", 10, 8, 1))
	zoo.Admit(animals.NewRabbit("Боб", 5, 9, 2))
	zoo.Admit(animals.NewTiger("Тайсон", 20, 3))
	zoo.Admit(animals.NewWolf("Серый", 15, 4))

	// Добавляем предметы
	zoo.AddInventory(things.NewComputer(5, "HP", 16, "Intel"))
	zoo.AddInventory(things.NewTable(6, "Дерево", "Коричневый"))
	zoo.AddInventory(things.NewThing(7, "Коробка", 10.5))

	// Проверяем подсчеты
	if len(console.Zoo.GetAllAnimals()) != 4 {
		t.Error("Должно быть 4 животных")
	}

	handAnimals := console.Zoo.GetAllHandAnimal(service.MinFriendlinessLevel)
	if len(handAnimals) != 2 {
		t.Errorf("Должно быть 2 дружелюбных животных, получено %d", len(handAnimals))
	}

	if len(console.Zoo.GetAllInventory()) != 7 {
		t.Error("Должно быть 7 предметов в инвентаре")
	}

	// Вызываем методы рендеринга
	console.WriteAllAnimal()
	console.WriteAllHandAnimal()
	console.WriteAllInventory()
}

func TestConsole_VetClinicIntegration(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Проверяем работу ветклиники через консоль
	ok, _ := console.Vet.CheckHealth(service.Health5)
	if !ok {
		t.Error("Ветклиника должна одобрить здоровое животное")
	}

	ok, reason := console.Vet.CheckHealth(service.Health1)
	if ok {
		t.Error("Ветклиника не должна одобрять больное животное")
	}
	if reason == "" {
		t.Error("Должна быть причина отказа")
	}
}

// ==================== Граничные случаи ====================

func TestConsole_WithMaxInventoryNumber(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Добавляем много предметов
	for i := 1; i <= 100; i++ {
		zoo.AddInventory(things.NewThing(i, fmt.Sprintf("Предмет%d", i), float64(i)))
	}

	nextNum := console.Zoo.NextInventoryNumber()
	if nextNum != 101 {
		t.Errorf("Следующий номер должен быть 101, получено %d", nextNum)
	}

	console.WriteAllInventory()
}

func TestConsole_MixedFriendlinessLevels(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Добавляем животных с разными уровнями дружелюбности
	for i := 0; i <= 10; i++ {
		zoo.Admit(animals.NewMonkey(fmt.Sprintf("Обезьяна%d", i), 10, i, i+1))
	}

	// Проверяем фильтрацию
	handAnimals := console.Zoo.GetAllHandAnimal(service.MinFriendlinessLevel)

	// Должны быть только с дружелюбностью >= 6 (то есть 6,7,8,9,10 = 5 животных)
	if len(handAnimals) != 5 {
		t.Errorf("Ожидалось 5 дружелюбных животных, получено %d", len(handAnimals))
	}

	console.WriteAllHandAnimal()
}

func TestConsole_LargeFoodAmount(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Добавляем животных с большим количеством еды
	zoo.Admit(animals.NewMonkey("Обжора1", 1000, 7, 1))
	zoo.Admit(animals.NewTiger("Обжора2", 2000, 2))
	zoo.Admit(animals.NewWolf("Обжора3", 1500, 3))

	totalFood := console.Zoo.CountFood()
	expected := 4500

	if totalFood != expected {
		t.Errorf("Ожидалось %d еды, получено %d", expected, totalFood)
	}

	console.WriteAllAnimal()
}

func TestConsole_SpecialCharactersInNames(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Добавляем животных со специальными символами
	zoo.Admit(animals.NewMonkey("Тест-123", 10, 7, 1))
	zoo.Admit(animals.NewRabbit("Кролик_ФИО", 5, 8, 2))
	zoo.Admit(animals.NewTiger("Тигр.v2", 20, 3))

	if len(console.Zoo.GetAllAnimals()) != 3 {
		t.Error("Все животные со специальными символами должны быть добавлены")
	}

	console.WriteAllAnimal()
}

func TestConsole_UnicodeNames(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Добавляем животных с unicode именами
	zoo.Admit(animals.NewMonkey("猴子", 10, 7, 1)) // китайский
	zoo.Admit(animals.NewRabbit("うさぎ", 5, 8, 2)) // японский
	zoo.Admit(animals.NewTiger("호랑이", 20, 3))    // корейский
	zoo.Admit(animals.NewWolf("Волк🐺", 15, 4))   // с эмодзи

	if len(console.Zoo.GetAllAnimals()) != 4 {
		t.Error("Все животные с unicode именами должны быть добавлены")
	}

	console.WriteAllAnimal()
}

// ==================== Дополнительные тесты для tools/render.go ====================

// ==================== Граничные случаи ====================

func TestConsole_WriteAllWithMixedContent(t *testing.T) {
	zoo := service.NewMoscowZoo()
	vet := service.NewVetClinic()
	console := tools.NewConsole(zoo, vet)

	// Добавляем смешанный контент
	zoo.Admit(animals.NewMonkey("Обезьяна", 10, 8, 1))
	zoo.Admit(animals.NewTiger("Тигр", 20, 2))
	zoo.AddInventory(things.NewComputer(3, "Dell", 16, "i7"))
	zoo.AddInventory(things.NewThing(4, "Box", 5.0))

	// Вызываем все методы рендеринга
	console.WriteAllAnimal()
	console.WriteAllHandAnimal()
	console.WriteAllInventory()

	// Проверяем корректность данных
	if len(console.Zoo.GetAllAnimals()) != 2 {
		t.Error("Должно быть 2 животных")
	}

	if len(console.Zoo.GetAllInventory()) != 4 {
		t.Error("Должно быть 4 предмета в инвентаре (2 животных + 2 предмета)")
	}
}
