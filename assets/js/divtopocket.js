//pocketmoney/assets/s/js
//requires shared js, replaceAll from template.js

function divtopocket(d,svg){
    if (!d.innerHTML) {
        console.log("No D innerhtml", d);
        return;
    }
    var s = d.innerHTML.split(":");
    if (s.length !== 2 ){
        s = ["black","red"];
    }
    if (s[0] == "" ){
        s[0] = "black";
    }
    if (s[1] ==""){
        s[1] = "red";
    }
    d.innerHTML = replaceAll(svg,{
        "maincol":s[0],
        "linecol":s[1]
    });
}

function divstopocket(svg){
    var ar = document.getElementsByClassName("pocket");
    for (var s in ar){
        divtopocket(ar[s],svg);
    }
}
function changecolor(){
        var c1 = document.getElementById("Col1");
        var c2 = document.getElementById("Col2");
        var pocket = document.getElementById("color_pocket");
        
        pocket.innerHTML = replaceAll(psvg,{maincol:c1.value,linecol:c2.value});
    }
