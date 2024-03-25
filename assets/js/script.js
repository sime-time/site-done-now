// HAMBURGER 
const hamburger = document.querySelector(".hamburger"); 
const navMenu = document.querySelector(".nav-menu");
const navLinks = document.querySelectorAll(".nav-link");

hamburger.addEventListener("click", () => {
    hamburger.classList.toggle("active");
    navMenu.classList.toggle("active");
});

navLinks.forEach(n => n.addEventListener("click", () => {
    hamburger.classList.remove("active");
    navMenu.classList.remove("active");
}));
 

// ACCORDION  
const accordionItems = document.querySelectorAll('.accordion-item');

accordionItems.forEach((item, index) => {
    let label = item.querySelector(".accordion-label");
    label.addEventListener("click", () => {

        item.classList.toggle("active");  
       
        let content = item.querySelector(".accordion-content"); 
        let icon = item.querySelector("i"); 

        if (item.classList.contains("active")) {

            content.style.paddingBottom = "2rem"
            content.style.height = `${content.scrollHeight}px`;
            
            icon.style.transform = "rotate(45deg)";
            icon.style.color = "var(--secondary-color)"; 

        } else {

            content.style.paddingBottom = "0"
            content.style.height = "0px";

            icon.style.transform = "rotate(0deg)";
            icon.style.color = "inherit"; 
        }
        
        // remove active status on all other items 
        removeActive(index); 
    });
});


function removeActive(index) {
    accordionItems.forEach((itm, idx) => {
        if (index != idx) {
            itm.classList.remove("active"); 

            let cnt = itm.querySelector(".accordion-content"); 
            cnt.style.paddingBottom = "0"; 
            cnt.style.height = "0px"; 

            let i = itm.querySelector("i"); 
            i.style.transform = "rotate(0deg)";
            i.style.color = "inherit";
        }
    })
}
