const faBars=document.querySelector('.fa-bars');
const contentLink=document.querySelector('.content-link');
const  btnUpdate=document.querySelector('.btn-update');
const formUpdate=document.querySelector('.form-update');
const faxmark=document.querySelector('.faxmark');
const  messageErro=document.querySelectorAll('.messageErro');
const errorPost=document.querySelector('.error-post');
const alertLikePost=document.querySelector('.alert-likePost');
const likeDeconnected=document.querySelectorAll('.like-deconnected');
const signin=document.querySelector('.signin');
const register=document.querySelector('.register');
// console.log(errorPost);
window.addEventListener('load',()=>{
       let loc=window.location.href;
       if (signin && loc.includes('signin')) {
            signin.style.border=1+"px  solid var(--header-color)"; 
            signin.style.padding= 7+"px "+15+"px";
       }
       if (register && loc.includes('register')) {
            register.style.border=1+"px  solid var(--header-color)"; 
            register.style.padding= 7+"px "+15+"px";
        }
        if (alertLikePost) {
            alertLikePost.style.display='none';
        }    
        // display the error post form
        if (errorPost) {
            let childsErroPost=errorPost.children
            if (childsErroPost[1].textContent.trim()=="") {
                childsErroPost[0].style.display="none";
            }
        }

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
                    }, 2000);
                    
                 }                                    
            }
           
        });
    }
})

if (alertLikePost && likeDeconnected) {
    for (let index = 0; index < likeDeconnected.length; index++) {
        const btn = likeDeconnected[index];
        btn.addEventListener("click",()=>{
            alertLikePost.style.display='block';
        })
    }
}