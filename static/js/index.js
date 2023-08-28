const btnNewPost=document.querySelector('.new-post');
const formNewPost=document.querySelector('.form-new-post');
const btnFilter=document.querySelector('.btn-filter');
const filterForm=document.querySelector('.filter-form');
const disabledBtn=document.querySelector('.disabled-btn');
const postDisplaying=document.querySelector(".post-displaying").children;
// pagination button
const  pagination=document.querySelector('.pagination');

var paginationCompt=Math.round(postDisplaying.length/10)


window.onload = () => {
     console.log(paginationCompt);
     if (paginationCompt==0) {
        pagination.style.display="none"
     }
     for (let j = 1; j <= paginationCompt; j++) {       
        const a=document.createElement('a');
        a.textContent=j
        a.classList=`page page${j}`
        pagination.appendChild(a)
     }
     const aa=document.createElement('a');
      aa.textContent="\""
      pagination.appendChild(aa)
 
     const pages=document.querySelectorAll(".page")
     console.log(pages);
    if (postDisplaying.length>10) {   
       var bool=false      


       for (let i = 1; i <= pages.length; i++) {
         const pa =document.querySelector(`.page${i}`);
         console.log(pa);
         pa.addEventListener("click", () => {
            
            for (let index = 0; index < postDisplaying.length; index++) {
                const categorieBloc = postDisplaying[index];
                if (index>postDisplaying.length/pages.length) {
                    bool=true
                    categorieBloc.style.display="none";
                } else{
                    categorieBloc.style.display="block";
                }
                
                
             }
         });
        
       }
   
         if (!bool){
            for (let index = 0; index < postDisplaying.length; index++) {
                const categorieBloc = postDisplaying[index];
                if (index>=10) {
                    categorieBloc.style.display="none";
                } 
                
             }   
         }
    }



    if (formNewPost && btnNewPost) {
        formNewPost.style.display="none";
        btnNewPost.addEventListener("click", (event) => {
            event.preventDefault();
            if (formNewPost.style.display === "none") {
                formNewPost.style.display = "block";
                btnNewPost.innerHTML="<i class='fa-solid fa-xmark'></i> Fermer"
            } else {
                btnNewPost.innerHTML="<i class='fa-solid fa-plus'></i> New Post"
                formNewPost.style.display = "none";
            }
        });
    }
   
    filterForm.style.display="none";

    btnFilter.addEventListener("click", (event) => {
        event.preventDefault();
        if (filterForm.style.display === "none") {
            filterForm.style.display = "block";
            btnFilter.innerHTML="<i class='fa-solid fa-xmark'></i> Fermer"
        } else {
            btnFilter.innerHTML="<i class='fa-solid fa-filter'></i> filter"
            filterForm.style.display = "none";
        }
    });
    if (disabledBtn) {
        disabledBtn.setAttribute("disabled", '');
    }
        
  
  
  };


