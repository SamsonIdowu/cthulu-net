package bot

import "log"

type Bot struct {
	uuid   string
	system System
}

// @notice Allows the Bot to ping the System and let it know the bot is live.
// @notice The bot receives in exchange a uuid for accessing tasks later.
// @param `tracker Tracker`, a proxy to the System, implementing the Tracker interface.
func (b *Bot) Ping(tracker Tracker) (err error){
	uuid, err := tracker.add(b)
	if err != nil {
		//panic(err)
		log.Println(err)
		return err
	}
	if len(uuid) > 40 {
		return
	}
	if b.uuid == "" {
		b.uuid = uuid
	}
	log.Printf("%v\n", b)
	return nil
}

func (b *Bot) Work(recipe Recipe, task Task) Result {
	if task.Id == "" {
		return Result{}
	}
	return recipe.Do(task)
}

func (b *Bot) Getuuid() (string){
	return b.uuid
}

func CreateBotInstance(sys System) Bot {
	return Bot{uuid: "", system: sys}
}
