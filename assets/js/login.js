//--- login.js ---
function shownewfam(){
	fb = document.getElementById("newfamily");
	lb = document.getElementById("login");
	fb.style.display = "";
	lb.style.display = "none";
	document.getElementById("famtop").focus();

}
function showlogin(f,u){
	var fb = document.getElementById("newfamily");
	var lb = document.getElementById("login");
	fb.style.display = "none";
	lb.style.display = "";

		
	var ftxt = document.getElementById("linfam");
	var lpass = document.getElementById("linpass");
	lpass.value = "";
	ftxt.focus();
	if (f !== undefined ){
		ftxt.value = f;
	}

	if (u !== undefined) {
		var utxt = document.getElementById("linusr");
		utxt.value = u;
		
		lpass.focus();
	}
}

