package errorlist

import (
	"errors"
	"fmt"
)

var (
	ErrAccessDenied   = errors.New("Процесс остановлен")
	ErrNotExist       = errors.New("Файл не существует")
	ErrForbiddenName  = errors.New("Запрещённое имя")
	ErrForbiddenChars = errors.New("Запрещённые символы")

	ErrNoSubEvent     = errors.New("SubEvent не задан")
	ErrUnknownCommand = errors.New("Неизвестная команда")

	ErrEmptyCommand = errors.New("Пустая команда")
	ErrTooFewArgs   = errors.New("Недостаточно аргументов")
	ErrNoPayload    = errors.New("Payload не найдена")

	ErrorInwork = errors.New("В разработке!")

	ErrTempDirEmpty = errors.New("os.TempDir() вернул пустой путь")
	ErrInitRoot     = errors.New("Ошибка инициализации!")

	ErrScriptOff = errors.New("Отключена работа со сценариями -sct")
)

func wrapError(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func InvalidName(name string, err error) error {
	return wrapError(fmt.Sprintf("Недопустимое имя %q", name), err)
}

func Create(name string, err error) error {
	return wrapError(fmt.Sprintf("Создание %s", name), err)
}

func Run(name string, err error) error {
	return wrapError(fmt.Sprintf("Запуск %s", name), err)
}

func Remove(name string, err error) error {
	return wrapError(fmt.Sprintf("Удаление %s", name), err)
}

func Read(name string, err error) error {
	return wrapError(fmt.Sprintf("Чтение %s", name), err)
}

func Open(name string, err error) error {
	return wrapError(fmt.Sprintf("Открытие %s", name), err)
}

func ClearLog(err error) error {
	return wrapError("Очистка журнала", err)
}

func RunCommand(command string, err error) error {
	return wrapError(fmt.Sprintf("Выполнение команды %q", command), err)
}

func ClearTerminal(err error) error {
	return wrapError("Очистка терминала", err)
}

func OpenFile(path string, err error) error {
	return wrapError(fmt.Sprintf("Открытие %s", path), err)
}

func ReadFile(path string, err error) error {
	return wrapError(fmt.Sprintf("Чтение %s", path), err)
}

func CreateFile(path string, err error) error {
	return wrapError(fmt.Sprintf("Создание %s; возможно, требуется -init", path), err)
}

func WriteFile(path string, err error) error {
	return wrapError(fmt.Sprintf("Запись %s", path), err)
}

func RemoveFile(path string, err error) error {
	return wrapError(fmt.Sprintf("Удаление %s", path), err)
}
