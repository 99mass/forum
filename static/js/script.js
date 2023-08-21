const faBars=document.querySelector('.fa-bars');
const contentLink=document.querySelector('.content-link');
// console.log(contentLink);

window.onload = () => {
    contentLink.style.display="none";
    faBars.addEventListener("click", (event) => {
        
        if (contentLink.style.display === "none") {
            contentLink.style.display = "block";
        } else {
            contentLink.style.display = "none";
        }
    });
};