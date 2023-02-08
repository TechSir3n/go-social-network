
if (window.location.href == "http://localhost:8000/login") {
    const button_b1 = document.getElementById('buttonID');

    button_b1.onclick = function GetUserName() {
        let getName = document.getElementById('nameID').value;
        localStorage.setItem('name', getName);
        console.log(localStorage.getItem('name'));
    }
}


if (window.location.href == "http://localhost:8000/access/admin") {
    const button_b2 = document.getElementById('buttonADMIN');

    button_b2.onclick = function GetAdminName() {
        let getName = document.getElementById('name_admin').value;
        localStorage.setItem('name_admin', getName);
        console.log(localStorage.getItem('name_admin'));
    }

}

