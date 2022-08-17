async function example(){
    document.getElementById("example").append(await rp_getENV("OS"))
}

example()