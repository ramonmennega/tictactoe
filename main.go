package main

import (
	"encoding/json"
	"fmt"
	//"encoding/json"
	//
	//"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var KKK = Tictactoe{
	Beurt: "1",
Win:   "",
P1: 0,
P2: 0,
Username1: "",
Username2: "",
Bord: Speelveld{
[9]string{"", "", "", "", "", "", "", "", ""},
},
}
var won string
var NOC int
var connections map[*websocket.Conn]bool
var connected = make(map[*websocket.Conn]bool)
var spelercounter int

type Speelveld struct {
	Veld [9]string
}

type Tictactoe struct {
	Beurt string
	Win   string
	P1	  int
	P2	  int
	Username1 string
	Username2 string
	Bord  Speelveld
}

func game(conn *websocket.Conn) {

	KKKK := KKK

	for k, _ := range connected {
		Test2, _ := json.Marshal(KKKK)
		k.WriteMessage(1, Test2)
	}

	for {
		_, msg, err := conn.ReadMessage()
		birdJson := string(msg)
		var obj Tictactoe
		json.Unmarshal([]byte(birdJson), &obj)
		if obj.Beurt != ""{
			if obj.Beurt == "1"{
				fmt.Printf("er is een zet gedaan door: speler 2\n")
			} else if obj.Beurt == "2"{
				fmt.Printf("er is een zet gedaan door: speler 1\n")
			}
		}
		if obj.Bord.Veld[0] == obj.Bord.Veld[1] && obj.Bord.Veld[0] == obj.Bord.Veld[2] && obj.Bord.Veld[0] == "X" { // kijkt wie gewonnen heeft
			obj.Win = "X"
		} else if obj.Bord.Veld[3] == obj.Bord.Veld[4] && obj.Bord.Veld[3] == obj.Bord.Veld[5] && obj.Bord.Veld[3] == "X" {
			obj.Win = "X"
		} else if obj.Bord.Veld[6] == obj.Bord.Veld[7] && obj.Bord.Veld[6] == obj.Bord.Veld[8] && obj.Bord.Veld[6] == "X" {
			obj.Win = "X"
		} else if obj.Bord.Veld[0] == obj.Bord.Veld[3] && obj.Bord.Veld[0] == obj.Bord.Veld[6] && obj.Bord.Veld[0] == "X" {
			obj.Win = "X"
		} else if obj.Bord.Veld[1] == obj.Bord.Veld[4] && obj.Bord.Veld[1] == obj.Bord.Veld[7] && obj.Bord.Veld[1] == "X" {
			obj.Win = "X"
		} else if obj.Bord.Veld[2] == obj.Bord.Veld[5] && obj.Bord.Veld[2] == obj.Bord.Veld[8] && obj.Bord.Veld[2] == "X" {
			obj.Win = "X"
		} else if obj.Bord.Veld[0] == obj.Bord.Veld[4] && obj.Bord.Veld[0] == obj.Bord.Veld[8] && obj.Bord.Veld[0] == "X" {
			obj.Win = "X"
		} else if obj.Bord.Veld[2] == obj.Bord.Veld[4] && obj.Bord.Veld[2] == obj.Bord.Veld[6] && obj.Bord.Veld[2] == "X" {
			obj.Win = "X"
			//	nu begint o
		} else if obj.Bord.Veld[0] == obj.Bord.Veld[1] && obj.Bord.Veld[0] == obj.Bord.Veld[2] && obj.Bord.Veld[0] == "O" {
			obj.Win = "O"
		} else if obj.Bord.Veld[3] == obj.Bord.Veld[4] && obj.Bord.Veld[3] == obj.Bord.Veld[5] && obj.Bord.Veld[3] == "O" {
			obj.Win = "O"
		} else if obj.Bord.Veld[6] == obj.Bord.Veld[7] && obj.Bord.Veld[6] == obj.Bord.Veld[8] && obj.Bord.Veld[6] == "O" {
			obj.Win = "O"
		} else if obj.Bord.Veld[0] == obj.Bord.Veld[3] && obj.Bord.Veld[0] == obj.Bord.Veld[6] && obj.Bord.Veld[0] == "O" {
			obj.Win = "O"
		} else if obj.Bord.Veld[1] == obj.Bord.Veld[4] && obj.Bord.Veld[1] == obj.Bord.Veld[7] && obj.Bord.Veld[1] == "O" {
			obj.Win = "O"
		} else if obj.Bord.Veld[2] == obj.Bord.Veld[5] && obj.Bord.Veld[2] == obj.Bord.Veld[8] && obj.Bord.Veld[2] == "O" {
			obj.Win = "O"
		} else if obj.Bord.Veld[0] == obj.Bord.Veld[4] && obj.Bord.Veld[0] == obj.Bord.Veld[8] && obj.Bord.Veld[0] == "O" {
			obj.Win = "O"
		} else if obj.Bord.Veld[2] == obj.Bord.Veld[4] && obj.Bord.Veld[2] == obj.Bord.Veld[6] && obj.Bord.Veld[2] == "O" {
			obj.Win = "O"
		} else if obj.Bord.Veld[0] == "" && obj.Bord.Veld[1] == "" && obj.Bord.Veld[2] == "" && obj.Bord.Veld[3] == "" && obj.Bord.Veld[4] == "" && obj.Bord.Veld[5] == "" && obj.Bord.Veld[6] == "" && obj.Bord.Veld[7] == "" && obj.Bord.Veld[8] == "" {
			obj.Win = ""
			fmt.Printf("de game is gerestart\n")
		}

		if obj.Win == "O"{	// set de winaar in de console
			fmt.Printf("Er is gewonnen! De winnaar is Speler 2\n");
			obj.P2++

			logWin("O")

		} else if obj.Win == "X"{
			fmt.Printf("Er is gewonnen! De winnaar is Speler 1\n");
			logWin("X")
			obj.P1++

		}



		msg2, _ := json.Marshal(obj)
		for k, _ := range connected {
			k.WriteMessage(1, msg2)
		}
		if err != nil {
			return
		}
		fmt.Printf("de zet is met succes verspreid over de clients\n")
	}
}

func ws(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	NOC++
	if NOC == 1 {
		conn.WriteMessage(1, []byte("speler1"))
		fmt.Printf("Er is een speler gejoind of geleaved. aantal is nu 1\n")
	} else if NOC == 2{
		conn.WriteMessage(1, []byte("speler2"))
		fmt.Printf("Er is een speler gejoind. aantal is nu 2\n")

	}
	if NOC > 2 {
		conn.WriteMessage(1, []byte("0"))
		connected[conn] = false
		conn.Close()
		NOC--
	} else {
		connected[conn] = true
	}
	conn.SetCloseHandler(func(c int, t string) error {
		NOC--
		for k, _ := range connected {
			k.WriteMessage(1, []byte("restart"))
		}
		http.Redirect(w,r,"/index.html", http.StatusSeeOther)
		connected[conn] = false
		conn.Close()
		return nil
	})

	game(conn)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			return
		}
	}
}

func main() {
	fmt.Printf("!!!!!!!!!LEES DIT EERST!!!!!!!!!\n\nClients kunnen joinen door de host ip in hun adressbalk te plaatsen met de port :9090 bijvoorbeeld 10.21.0.47:9090. het ip vind je via cmd en dan ipconfig.\n FaQ: \n\n Q: hoezo kan alleen ik spelen??\nA: Om samen te kunnen spelen moet je firewall instelling aanpassen en de tictactoe.exe toegang geven en openbaar maken.\n\n\n ")
	fmt.Printf("Programma is met succes gestart. \n\n")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/echo", ws)
	http.HandleFunc("/login", login)
	_ = http.ListenAndServe(":9090", nil)
}
