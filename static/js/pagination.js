const postDisplaying = document.querySelector(".post-displaying").children;
const pagination = document.querySelector('.pagination');
const paginationCompt = Math.ceil(postDisplaying.length / 10); // Utilisez Math.ceil pour arrondir à la hausse

console.log(postDisplaying.length);
if (paginationCompt <= 1) { // Modifié la condition ici
    pagination.style.display = "none";
}

for (let j = 1; j <= paginationCompt; j++) {
    const a = document.createElement('a');
    a.textContent = j;
    a.classList = `page page${j}`;
    pagination.appendChild(a);
}

// if (paginationCompt > 1) { // Modifié la condition ici
//     const aa = document.createElement('a');
//     aa.textContent = "»"; // Utilisez une flèche simple ici
//     aa.classList = `page-right`;
//     pagination.appendChild(aa);
// }

const pages = document.querySelectorAll(".page");
// const pageRight=document.querySelector(".page-right");
// const pageLeft=document.querySelector(".page-left");

// console.log(pagedroit);

if (postDisplaying.length > 10) {
    for (let i = 1; i <= pages.length; i++) {
        const pa = document.querySelector(`.page${i}`);

        pa.addEventListener("click", () => {

        });
        pa.addEventListener("click", () => {
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




