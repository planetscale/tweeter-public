package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gosimple/slug"
	"github.com/jaswdr/faker"
)

func main() {
	messages := []string{
		"Just had the most amazing cup of coffee ☕️ #MondayMotivation",
		"Excited to start my new job today! Wish me luck 🤞 #NewBeginnings",
		"Why is it that whenever I decide to go on a diet, all I crave is pizza? 🍕 #FoodieProblems",
		"Finally finished reading that book everyone's been talking about. Such a captivating story! 📚 #Bookworm",
		"Can't wait for the weekend to relax and binge-watch my favorite TV series. 📺 #NetflixAndChill",
		"Spent the day exploring a new hiking trail. Nature never fails to amaze me! 🌲 #OutdoorAdventure",
		"Currently stuck in traffic. Send help! 🚗 #RushHourProblems",
		"Just booked my tickets for a dream vacation! Counting down the days. ✈️ #Wanderlust",
		"It's amazing how a simple smile can brighten someone's day. 😊 #SpreadPositivity",
		"Trying out a new recipe tonight. Fingers crossed it turns out delicious! 🍽️ #HomeCooking",
		"Nothing beats a good workout to kickstart the day. Feeling energized! 💪 #FitnessGoals",
		"The best part of traveling is trying new cuisines. Today's food adventure: sushi! 🍣 #FoodExplorer",
		"Happiness is finding the perfect playlist that matches your mood. 🎶 #MusicLover",
		"Weekend plans: pajamas, popcorn, and a movie marathon. Pure bliss! 🍿 #LazyDays",
		"Just witnessed the most breathtaking sunset. Nature's artwork never ceases to amaze me. 🌅 #BeautyOfNature",
		"Why do Mondays always feel like a mission impossible? 😴 #MondayBlues",
		"Spent the entire day at the beach. Sand, sun, and good company—perfection! 🏖️ #BeachLife",
		"Just got tickets to see my favorite band live! This is going to be epic! 🎤 #MusicFestival",
		"Woke up to find my plants in full bloom. Gardening success! 🌺 #GreenThumb",
		"Rainy days are the perfect excuse to stay in bed with a good book. 🌧️ #CozyVibes",
		"Nothing feels better than crossing off all the tasks on my to-do list. Productivity level: 💯 #GettingThingsDone",
		"Attended an inspiring TED Talk today. So many great ideas worth spreading! 🎙️ #TEDx",
		"Just tried a new workout class and my muscles are already feeling the burn. Ouch! 💦 #FitnessJourney",
		"Celebrating a major milestone today! Hard work pays off. 🎉 #AchievementUnlocked",
		"Spent the day volunteering at a local shelter. It's amazing how giving back can make you feel so fulfilled. 🤝 #PayItForward",
		"The best part of traveling is making new friends from around the world. 🌍 #GlobalConnections",
		"Had the most delicious ice cream today. I think I found my new favorite flavor! 🍦 #SweetTooth",
		"Sunday mornings are made for pancakes and lazy cuddles. ❤️ #WeekendVibes",
		"Just finished an intense yoga session. Mind, body, and soul feel rejuvenated. 🧘‍♀️ #YogaLife",
		"Learning a new language is challenging but incredibly rewarding. Can't wait to become fluent! 🗣️ #LanguageLover",
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?tls=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
	}
	log.Println("Successfully connected to PlanetScale!")

	fake := faker.New()

	miter := 0

	for i := 0; i < 10; i++ {
		p := fake.Person()
		s := slug.Make(p.FirstName())

		query := `insert into users (name, email, slug, password_digest, created_at, updated_at)
			values (? , ?, ?, 'password123', ?, ?)`
		results, err := db.Exec(query,
			p.Name(), p.Contact().Email, s, getDateTime(fake), getDateTime(fake))
		if err != nil {
			log.Fatal(err)
		}
		iid, err := results.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		query2 := `insert into tweets(content, user_id, created_at, updated_at)
			values (?, ?, ?, ?)`
		for j := 0; j < 3; j++ {
			results, err = db.Exec(query2,
				messages[miter], iid, getDateTime(fake), getDateTime(fake))
			if err != nil {
				log.Fatal(err)
			}
			miter++
		}
	}
}

func getDateTime(f faker.Faker) string {
	nowUnix := time.Now().Unix()
	history := time.Now().Add(time.Hour * -1).Unix()
	unixtimestamp := rand.Int63n(nowUnix-history) + history

	return time.Unix(unixtimestamp, 0).Add(time.Hour * 4).Format("2006-01-02 15:04:05")
}
