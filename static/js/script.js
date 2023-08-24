const faBars=document.querySelector('.fa-bars');
const contentLink=document.querySelector('.content-link');
const  btnUpdate=document.querySelector('.btn-update');
const formUpdate=document.querySelector('.form-update');
const faxmark=document.querySelector('.faxmark');
    contentLink.style.display="none";
    faBars.addEventListener("click", () => {
        
        if (contentLink.style.display === "none") {
            contentLink.style.display = "block";
        } else {
            contentLink.style.display = "none";
        }
    });
if (formUpdate && btnUpdate) {
    formUpdate.style.display="none";
    btnUpdate.addEventListener("click", () => {
        
        if (formUpdate.style.display === "none") {
            formUpdate.style.display = "block";
        } else {
            formUpdate.style.display = "none";
        }
    });
    faxmark.addEventListener("click", () => {
        formUpdate.style.display = "none";
    });
};