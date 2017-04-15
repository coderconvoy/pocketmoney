//requires shared js, replaceAll from template.js
//

function divtopocket(d,svg){
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
