const faBars=document.querySelector('.fa-bars');
const contentLink=document.querySelector('.content-link');
const  btnUpdate=document.querySelector('.btn-update');
const formUpdate=document.querySelector('.form-update');
const faxmark=document.querySelector('.faxmark');
const  messageErro=document.querySelectorAll('.messageErro');
// const monform=document.querySelector('.monform');


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

if (messageErro.length>0) {
    window.addEventListener("load", () => {
        for (let index = 0; index < messageErro.length; index++) {
            const element = messageErro[index];
             if (element.textContent!="") {
                setTimeout(() => {
                    element.innerHTML=""
                }, 1500);
                
             }
                
            
        }
    });
}