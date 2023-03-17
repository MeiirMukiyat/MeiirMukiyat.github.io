let sections = document.querySelectorAll('section')
let navLinks = document.querySelectorAll('header nav a')

window.onscroll = () =>{
    sections.forEach(sec =>{
        let top =window.scrollY;
        let offset = sec.offsetTop -150;
        let height = sec.offsetHeight;
        let id = sec.getAttribute('id');

        if(top >= offset && top< offset+height){
            navLinks.forEach(links =>{
                links.classList.remove('active');
                document.querySelector('header nav  a[href*='+ id+']').classList.add('active');
            });
        };
    });
    let header = document.querySelector('header')

    header.classList.toggle('sticky',window.scrollY>100);
};



document.getElementById("mess_send").addEventListener("click", function() {
    let name = document.getElementById("name").value;
    let email = document.getElementById("email").value;
    let message = document.getElementById("message").value;
    let number = document.getElementById("number").value;

    if (!name || !email || !message) {
      alert("Please fill in all required fields");
    } else {
      let data = "Name: " + name + "\nEmail: " + email + "\nMessage: " + message + "\n\n";

      let blob = new Blob([data], { type: "text/plain" });

      let link = document.createElement("a");
      link.download = "message.txt";
      link.href = URL.createObjectURL(blob);
      link.click();
      showSuccessMessage();
    }
  });

  function showSuccessMessage() {
    let successMessage = document.querySelector(".success-message");
    let closeIcon = successMessage.querySelector(".close-icon");

    successMessage.classList.remove("hidden");
    closeIcon.addEventListener("click", function() {
      successMessage.classList.add("hidden");
      });
}  