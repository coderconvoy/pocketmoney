

//populate login page with famname,username
function logpopulate(fn,un){
    ftxt = document.getElementById("linfam");
    ftxt.value = fn;

    utxt = document.getElementById("linusr");
    utxt.value = un;

    document.getElementById("linpass").focus();


}
