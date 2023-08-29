
const postDisplaying=document.querySelector(".post-displaying").children;
// pagination button
const  pagination=document.querySelector('.pagination');

var paginationCompt=Math.round(postDisplaying.length/10)


     console.log(postDisplaying.length);
     if (paginationCompt==0) {
        pagination.style.display="none"
     }
     for (let j = 1; j <= paginationCompt; j++) {       
        const a=document.createElement('a');
        a.textContent=j
        a.classList=`page page${j}`
        pagination.appendChild(a)
     }
     if (pagination) {
        const aa=document.createElement('a');
            aa.textContent="\""
            pagination.appendChild(aa)
     }
     
 
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



