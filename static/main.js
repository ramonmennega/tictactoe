let speler = "";
let winner = false;
let output = document.getElementById("output");
let socket = new WebSocket("ws://10.21.0.47:9090/echo");
let obj = "";
let classname = document.getElementsByClassName("vakjes");
socket.onopen = function () {
    for (var i = 0; i < classname.length; i++) {
        classname[i].addEventListener('click', game);
    }
};

socket.onmessage = function (e) {
    if (e.data == "restart") {
        location.reload();
    }
    if (e.data == "speler1") {
        output.innerHTML += "U bent speler 1";

    } else if (e.data == "speler2") {
        output.innerHTML += "U bent speler 2";
        for (var i = 0; i < classname.length; i++) {
            classname[i].removeEventListener('click', game);
        }
    } else if (e.data == "0") {
        window.open("http://www.webdesign-forum.nl");

    }
    obj = JSON.parse(e.data);
    document.getElementById("speler1").innerHTML = obj.Username1;
    document.getElementById("speler2").innerHTML = obj.Username2;
    document.getElementById("score1"). innerHTML = obj.P1;
    document.getElementById("score2"). innerHTML = obj.P2;
    if (output.innerHTML == "U bent speler 1") {
        if (obj.Beurt == "2") {
            document.getElementById("beurt").innerHTML = obj.Username2 + " is aan de beurt";
            for (var i = 0; i < classname.length; i++) {
                classname[i].removeEventListener('click', game);
            }
        } else if (obj.Beurt == "1") {
            document.getElementById("beurt").innerHTML = obj.Username1 +  " is aan de beurt";

            for (var i = 0; i < classname.length; i++) {
                classname[i].addEventListener('click', game);
            }
        }
    } else if (output.innerHTML == "U bent speler 2") {
        if (obj.Beurt == "1") {
            document.getElementById("beurt").innerHTML = "Speler 1 is aan de beurt";

            for (var i = 0; i < classname.length; i++) {
                classname[i].removeEventListener('click', game);
            }
        } else if (obj.Beurt == "2") {
            document.getElementById("beurt").innerHTML = "Speler 2 is aan de beurt";

            for (var i = 0; i < classname.length; i++) {
                classname[i].addEventListener('click', game);
            }
        }
    }

    if (obj.Win == "X") {
        document.getElementById("winnaar").innerHTML = obj.Username1 + " heeft gewonnen!!";
    } else if (obj.Win == "O") {
        document.getElementById("winnaar").innerHTML = obj.Username2 + " heeft gewonnen!!";

    }
    if (obj.Beurt == 1) {
        beurtje = "X";
        obj.Beurt = "2";
    } else {
        beurtje = "O"
        obj.Beurt = "1";

    }
    speler = beurtje;
    vakarray();

};


function vakarray() {
    for (let i = 0; i < obj.Bord.Veld.length; i++) {

        if (obj.Bord.Veld[i] == "X") {

            document.getElementById(i).innerHTML = "X";

        } else if (obj.Bord.Veld[i] == "O") {

            document.getElementById(i).innerHTML = "O";
        } else if (obj.Bord.Veld[i] == "") {

            document.getElementById(i).innerHTML = "  ";
        }
    }


}

function game() {
    var vakje = event.target;

    var data = vakje.innerHTML;
    if (winner == false) {
        if (data == "  ") {
            vakje.innerHTML = speler;

        } else {
        }
    }

    function rolwissel(vakinhoud) {
        if (vakinhoud === "X") {
            obj.Bord.Veld[event.target.id] = "X";
            vakinhoud = "X";


        } else {
            obj.Bord.Veld[event.target.id] = "O";


            vakinhoud = "O"

        }

        send();

        return vakinhoud;

    }


    if (data == "  " && winner != true) {
        speler = rolwissel(speler);
        // bericht1(speler);
    }

}

function restartgame() {
    obj.Bord.Veld = ["", "", "", "", "", "", "", "", ""];
    if (obj.Beurt == 1){
        obj.Beurt = 2;
    } else if(obj.Beurt == 2){
        obj.Beurt = 1;
    }
    send();
}

function new_game() {
}

function send() {


    var test = JSON.stringify(obj);
    socket.send(test);

}