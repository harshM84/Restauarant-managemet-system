const txt1= document.getElementById('tbuser');/* tbuser is id name  */
const check= document.getElementById('check');
const outputt=document.getElementById('outputt');
const customername=document.getElementById("customername")
function fun1() {
    outputt.innerHTML = txt1.value;
}
check.addEventListener('click',fun1)
// <--------------------------script for check box ------------------------------->

// // Get the element to display selected items
// const selectedList = document.getElementById('selectedList');
// // Get all the checkboxes
// const checkboxes = document.querySelectorAll('input[type="checkbox"]');
//     function updateSelectedItems() {
//         // Clear the current list
//         selectedList.innerHTML = '';
//         // Get all checked checkboxes and add them to the list
//         checkboxes.forEach(checkbox => {
//             if (checkbox.checked) {
//                 const listItem = document.createElement('li');
//                 listItem.textContent = checkbox.value;
//                 selectedList.appendChild(listItem);
//             }
//         });
//     }
// // Add event listeners to all checkboxes
// checkboxes.forEach(checkbox => {
//         checkbox.addEventListener('change', updateSelectedItems);
//     });
// // Initial update to display any pre-checked items
// updateSelectedItems();