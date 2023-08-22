const faBars=document.querySelector('.fa-bars');
const contentLink=document.querySelector('.content-link');
const  btnUpdate=document.querySelector('.btn-update');
const formUpdate=document.querySelector('.form-update');
    contentLink.style.display="none";
    faBars.addEventListener("click", () => {
        
        if (contentLink.style.display === "none") {
            contentLink.style.display = "block";
        } else {
            contentLink.style.display = "none";
        }
    });
    formUpdate.style.display="none";
    btnUpdate.addEventListener("click", () => {
        
        if (formUpdate.style.display === "none") {
            formUpdate.style.display = "block";
        } else {
            formUpdate.style.display = "none";
        }
    });