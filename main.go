package main

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

/*
	create for interfaces sections
*/

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spinishBot struct{}

func main() {

	/* here is simple project 1 */

	// cards := newDeck()
	// cards.suffler()
	// cards.print()

	/*
		ตอนไหนถึงต้องใช้ interfaces จังหวะที่เรามีการกระทำบ้างอย่างที่คล้าย ๆ กันแต่อาจจะแตกต่างกันที่ ไทป์ของข้อมูล
		จำลองสถานะการณ์เรามี functions สำหรับแสดงข้อความต่างภาษาแต่มีการทำงานคล้ายกัน ต่างกันที่ไทป์ของตัวแปร
	*/

	/*	**************  bad ways ************** */

	/*
			เราไม่สามารถ overload methods ได้ใน Go
			การเขียนแบบนี้่จะทำให้่เราไม่สามารถยุบ methods ได้
			และไม่สามารถ complie ได้ถึงแม้จะเปลี่บยนชื่อแต่ก็ไม่สามารถทำให้ code มัน clean ได้


		eb := englishBot{}
		sb := spinishBot{}

		printGreeting(eb) // error
		printGreeting(sb) // error


		func (englishBot) getGreeting() string {
			return "hi there!"
		}

		func (spinishBot) getGreeting() string {
			return "hola !"
		}

		func printGreeting(eb englishBot) {
			fmt.Println(eb.getGreeting())
		}

		func printGreeting(sb spinishBot) {
			fmt.Println(sb.getGreeting())
		}


	*/

	/* **************  welcome interfaces ************** */

	/*
		การประกาศ interfaces

		type ชื่อไทป์ interface {
			methods()
		}

		ตัวอย่าง

		1. ประกาศ interfaces พร้อมกับ functions ที่เราต้องการ

			type bot interface {
				getGreeting() string // ต้องตั้งชื่อให้เหมือนกัน methods ของแต่ละ custom type ด้วย
			}


			func printGreeting(eb englishBot) {
				fmt.Println(eb.getGreeting())
			}

			func printGreeting(sb spinishBot) {
				fmt.Println(sb.getGreeting())
			}

		2. ทำประสร้าง method ที่เราต้องการโดยรับ พารามิเตอร์เป็นไทป์ของ interface

			func printGreeting(b bot) {
				fmt.println(b.getGreeting())
			}

		3. ทำการเรียกใช้ง่าน methods ที่เราสร้างขึ้นมา

			printGreeting(eb)
			printGreeting(sb)
	*/

	/*
		สรุป interface ใน Go คือโยนเข้ามาเลยไม่สนใจไทป์ ในแต่ละ type ต้องมี methods ที่ชื่อและ return type เหมือนกัน
		เพื่อเป็น association ของกันและกันจากในตัวอย่างคือ getGreeting() ที่คืนค่ากลับเป็น String นั้นคือ association
		ของ interface bot กับ englishBot , spiashBot

		ว่าง่าย ๆ ก็คือ interface เป็น set ของ function ของ custom type ที่มีการทำงานคล้าย ๆ กันใช้งานเพื่อลดความซ้ำซ้อนของ code 
	*/

}
