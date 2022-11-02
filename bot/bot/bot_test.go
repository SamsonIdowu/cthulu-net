package bot

import (
	"reflect"
	"testing"
)

func TestCreateBotInstance(t *testing.T) {
	t.Run("it should return a bot instance", func(t *testing.T) {
		// arrange
		ms := &MockSystem{}

		// act
		got := reflect.TypeOf(CreateBotInstance(ms))
		want := reflect.TypeOf(Bot{})

		// assert
		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

func TestCreateSystemBotInstance(t *testing.T) {
	t.Run("it should return a system bot instance", func(t *testing.T) {
		// arrange
		ms := &MockSystem{}
		ms.Init()

		// act
		want := Bot{uuid: "", system: ms}
		got := CreateBotInstance(ms)

		// assert
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

func TestBotCanPing(t *testing.T) {
	t.Run("it should ping tracker and receive uuid", func(t *testing.T) {
		// arrange
		tracker := MockTracker{}
		ms := &MockSystem{}
		bot := CreateBotInstance(ms)
		botptr := &bot
		// act
		botptr.Ping(tracker)
		// assert
		if botptr.uuid == "" {
			t.Errorf("failed to ping and get uuid")
		}
	})
}

func TestBotCanGetATask(t *testing.T) {
	t.Run("it should test if the bot can get a task object", func(t *testing.T) {

		// arrange
		tasker := MockTasker{}
		want := reflect.TypeOf(Task{})

		// act
		task := tasker.Next()
		got := reflect.TypeOf(task)
		// assert
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

func TestBotCanGetAParametizedTask(t *testing.T) {
	t.Run("it should be able to get a parametised task from Tasker Iface", func(t *testing.T) {
		// arrange
		want := Task{Id: "0", Ip: "10.1.1.4", Port: 8001, Type: "scan"}
		tasker := MockTasker{JsonObj: `{"id": "0", "ip": "10.1.1.4", "port": 8001, "type": "scan"}`}

		// act
		got := tasker.Next()

		// assert
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got : %v", got, want)
		}
	})

}

func TestBotShouldBeCapableOfReturningAResult(t *testing.T) {
	t.Run("it should return a result of task", func(t *testing.T) {
		// arrange
		task := Task{Id: "0", Ip: "10.1.1.4", Port: 8001, Type: "Scan"}
		want := Result{TaskId: "0", Ip: task.Ip, Port: int(task.Port), Ip_status: "up", Port_status: "filtered"}
		bot := CreateBotInstance(&MockSystem{})

		// act
		got := bot.Work(MockRecipe{}, task)

		// assert
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got : %v", got, want)
		}
	})
}
