

if (window.location.href == "http://localhost:8000/admin") {
    document.getElementById('id_h2').innerHTML = localStorage.getItem('name_admin');
} else {
    console.log('[ERROR:] Somethig went wrong,:if(inner_HTML)')
}



if (window.location.href == "http://localhost:8000/admin") {
    const button_add = document.getElementById('ADD_ID')
    button_add.onclick = function AddUser() {
        

    }

    const button_del = document.getElementById('DELETE_ID')
    button_del.onclick = function DeleteUser() {


    }

    const button_upd = document.getElementById('UPDATE_ID')
    button_upd.onclick = function UpdaeUser() {


    }

} else {
    console.log('[ERROR:] Somethig went wrong,:if(button_on_click)')
}