

async function e(){

    const dark = window.matchMedia('(prefers-color-scheme: dark)');

    if (dark.matches == true){
        document.body.setAttribute("mode", "dark")
    }else{
        document.body.setAttribute("mode", "light")
    }


}

e()

var timerID;
timerID = window.setInterval(e, 100);

async function example(){
    document.getElementById("example").append(await Env_GetVar("OS"))
}

example()
