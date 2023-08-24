const btnNewPost=document.querySelector('.new-post');
const formNewPost=document.querySelector('.form-new-post');
const btnFilter=document.querySelector('.btn-filter');
const filterForm=document.querySelector('.filter-form');
window.onload = () => {
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
  
  };


