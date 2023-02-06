const button = document.getElementById('buttonID');

button.onclick = function GetUserName() {
    let getName = document.getElementById('nameID').value;
    localStorage.setItem('name', getName);
    console.log(localStorage.getItem('name'));
}
