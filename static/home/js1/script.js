if (window.location.href == "http://localhost:8000/home") {
    document.getElementById('id_h').innerHTML = localStorage.getItem('name');
}


/*
if (window.location.href == "http://localhost:8000/home/update_password/password") {
    document.getElementById("btn_password_id").addEventListener("click", function () {
        var changePasswordDiv = document.querySelector(".change_password");
        changePasswordDiv.style.display = "block";
    });
}

if (window.location.href == "http://localhost:8000/home/update_email/email") {
    document.getElementById("btn_name_id").addEventListener("click", function () {
        var changePasswordDiv = document.querySelector(".change_email");
        changePasswordDiv.style.display = "block";
    });
}
*/

if (window.location.href == "http://localhost:8000/home/settings") {
    document.getElementById("btn_email_id").addEventListener("click", function () {
        var changePasswordDiv = document.querySelector(".change_name");
        changePasswordDiv.style.display = "block";
    });
}






