// JavaScript code for handling button click
document.addEventListener('DOMContentLoaded', function() {
    const imageButton = document.getElementById('imageButton');
    
    imageButton.addEventListener('click', function() {
        alert('Image button clicked!');
        
        // You can add more actions here, like navigating to a new page or performing some action.
    });
});

function myFunction() {
    var x = document.getElementById("myTopnav");
    if (x.className === "topnav") {
      x.className += " responsive";
    } else {
      x.className = "topnav";
    }
  }
function openNav() {
    document.getElementById("mySidebar").style.width = "250px";
    document.getElementById("main").style.marginLeft = "250px";
  }
  
function closeNav() {
    document.getElementById("mySidebar").style.width = "0";
    document.getElementById("main").style.marginLeft= "0";
  }

// -----------------scrpit for waiter.html--------------------

