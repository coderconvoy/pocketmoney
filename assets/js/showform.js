function showform(fid){
    var fdiv = document.getElementById("formlist");
    for( c in fdiv.children){
        console.log(fdiv.children[c]);
        if (fdiv.children[c].style) {
            fdiv.children[c].style.display = "none";
        }
    }
    document.getElementById(fid).style.display= "";
}
