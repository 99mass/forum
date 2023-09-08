const btnNewPost=document.querySelector('.new-post');
const formNewPost=document.querySelector('.form-new-post');
const btnFilter=document.querySelector('.btn-filter');
const filterForm=document.querySelector('.filter-form');
const btnLike=document.querySelector('.btn-like');
const disabledBtn=document.querySelector('.disabled-btn');
const  all=document.querySelector('.all');
const categoryId=document.querySelector('.category-id');
const categoriesId=document.querySelectorAll('.categories-id');
const submitBtn=document.querySelectorAll('.submit-btn');
const postDisplayingPostliked=document.querySelector('.post-displaying-postliked');
const postDisplayingDefault=document.querySelector('.post-displaying-default');
const alertPostLiked=document.querySelector('.alert-post-liked');

window.onload = () => {
    if (btnFilter && postDisplayingPostliked) {
        postDisplayingPostliked.style.display='none';
        btnLike.addEventListener('click',()=>{
            if (postDisplayingPostliked.style.display==='none') {
                postDisplayingDefault.style.display='none';
                postDisplayingPostliked.style.display='block';
                btnLike.innerHTML="<i class='fa-solid fa-xmark'></i> Fermer";
                if (postDisplayingPostliked.children.length>1) {
                    alertPostLiked.style.display='none';
                }
                pagination.style.display='none';
                
            }else{
                postDisplayingDefault.style.display='block';
                postDisplayingPostliked.style.display='none';
                btnLike.innerHTML="<i class='fa-solid fa-heart' style='margin-right: 5px;color: brown;'></i>Post-Liked";
            }
           
        });
    }
    
    
    // open or non open the post form
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
//    open or non the the filter form
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
    if (categoriesId.length>0 && categoryId) {
        let boole=false
        for (let index = 0; index < categoriesId.length; index++) {
            const cat = categoriesId[index];
            if (cat.value.trim()=== categoryId.textContent.trim()) {
                console.log(categoryId);
                console.log(cat.value);
                submitBtn[index].style.background="var(--color-linear-gradient)";
                submitBtn[index].style.color="#fff";
                submitBtn[index].style.border="none";
                submitBtn[index].style.borderRadius=0.5+"rem";
                boole=true
            }
            
        }
        if (!boole) {
            all.style.background="var(--color-linear-gradient)";
            all.style.color="#fff";
            all.style.border="none";
            all.style.borderRadius=0.5+"rem";
        }
    }
  
    // display the error post form
    if (errorPost) {
        let childsErroPost=errorPost.children
        if (childsErroPost[1].textContent.trim()=="") {
            childsErroPost[0].style.display="none";
        }else{
            formNewPost.style.display = "block";
            btnNewPost.innerHTML="";
            btnNewPost.innerHTML="<i class='fa-solid fa-xmark'></i> Fermer";
        }
    }



  };


