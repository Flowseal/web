function changePasswordVisibility() {
    let img = document.querySelector(".password-eye-icon");
    let input = document.getElementById("password");
    if (input.type == "password")
    {
        img.src = "static/img/eye-off.svg";
        input.type = "text";
    }
    else 
    {
        img.src = "static/img/eye-on.svg";
        input.type = "password";
    }
}