const faBars=document.querySelector('.fa-bars');
const contentLink=document.querySelector('.content-link');
const  btnUpdate=document.querySelector('.btn-update');
const formUpdate=document.querySelector('.form-update');
const faxmark=document.querySelector('.faxmark');
const errorPost=document.querySelector('.error-post');
const errorPostFilter=document.querySelector('.error-post-filter');
const alertLikePost=document.querySelector('.alert-likePost');
const likeDeconnected=document.querySelectorAll('.like-deconnected');
const signin=document.querySelector('.signin');
const register=document.querySelector('.register');

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
        if (errorPostFilter) {
            let childsErroPos=errorPostFilter.children
            if (childsErroPos[1].textContent.trim()=="") {
                childsErroPos[0].style.display="none";
            }else{
                filterForm.style.display = "block";
                btnFilter.innerHTML="<i class='fa-solid fa-xmark'></i> Fermer"
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
    if (alertLikePost && likeDeconnected) {
        for (let index = 0; index < likeDeconnected.length; index++) {
            const btn = likeDeconnected[index];
            btn.addEventListener("click",()=>{
                window.scrollTo(0, 0);
                alertLikePost.style.display='block';
            })
        }
    }
    const show=document.querySelector('.show');
    var x = document.getElementById("motdepasse");
    console.log(show);
    show.addEventListener('click',()=>{
        if (x.type === "password") {
            x.type = "text";
            show.innerHTML="<i class='fa-regular fa-eye' style='color:gray'></i>";
          } else {
            x.type = "password";
            show.innerHTML="<i class='fa-regular fa-eye-slash'></i>";
          }
    })

})

if (alertLikePost && likeDeconnected) {
    for (let index = 0; index < likeDeconnected.length; index++) {
        const btn = likeDeconnected[index];
        btn.addEventListener("click",()=>{
            window.scrollTo(0, 0);
            alertLikePost.style.display='block';
            setTimeout(() => {
                if (alertLikePost.style.display!=='none') {
                        alertLikePost.style.display='none';                
                }
            }, 5000);
        })
    }


}
