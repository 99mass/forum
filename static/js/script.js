const faBars=document.querySelector('.fa-bars');
const contentLink=document.querySelector('.content-link');


    contentLink.style.display="none";
    faBars.addEventListener("click", () => {
        
        if (contentLink.style.display === "none") {
            contentLink.style.display = "block";
        } else {
            contentLink.style.display = "none";
        }
    });