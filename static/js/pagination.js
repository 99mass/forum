const postDisplaying = document.querySelector(".post-displaying").children;
const pagination = document.querySelector('.pagination');
const paginationCompt = Math.ceil(postDisplaying.length / 10); 




for (let j = 1; j <= paginationCompt; j++) {
    const a = document.createElement('a');
    a.textContent = j;
    a.classList = `page page${j}`;
    pagination.appendChild(a);
}
if (paginationCompt <= 1) { 
    pagination.style.display = "none";
    
}
if ( postDisplaying.length<10) {
    document.querySelector('.page-page').style.display = "none";
}


const pages = document.querySelectorAll(".page");

if (postDisplaying.length > 10) {
    for (let i = 1; i <= pages.length; i++) {
        const pa = document.querySelector(`.page${i}`);

        pa.addEventListener("click", () => {
            window.scrollTo(0, 0);
            pages.forEach(page => page.classList.remove("actives"));
            pa.classList.add("actives");

            for (let index = 0; index < postDisplaying.length; index++) {
                const categorieBloc = postDisplaying[index];
                if (i !== Math.ceil((index + 1) / 10)) { 
                    categorieBloc.style.display = "none";
                } else {
                    categorieBloc.style.display = "block";
                }
            }
           
        });
    }

    for (let index = 0; index < postDisplaying.length; index++) {
        pages[0].classList="actives"
        const categorieBloc = postDisplaying[index];
        if (index >= 10) {
            categorieBloc.style.display = "none";
        }
    }
}




