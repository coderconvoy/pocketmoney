function showform(fid){
    var fdiv = document.getElementById("formlist");
    if (fdiv === undefined){
        console.Log("No FDiv for 'showform'");
        return
    }

    for( c in fdiv.children){
        console.log(fdiv.children[c]);
        if (fdiv.children[c].style) {
            fdiv.children[c].style.display = "none";
        }
    }
    if (fid === undefined){
        fdiv.children[0].style.display ="";
        return 
    }
    document.getElementById(fid).style.display= "";
}


