document.getElementById("uploadForm").addEventListener("submit", function (e) {
    e.preventDefault();

    var fileInput = document.getElementById("fileInput");
    var file = fileInput.files[0];

    var formData = new FormData();
    formData.append("file", file);

    fetch("/upload", {
        method: "POST",
        body: formData
    })
        .then(response => response.json())
        .then(data => {
            var fileList = document.getElementById("fileList");
            var listItem = document.createElement("li");

            var imgSrc = "http://localhost:8888/download/" + data.hash;
            listItem.innerText = file.name + " - " + data.hash;

            var img = document.createElement("img");
            img.src = imgSrc;

            listItem.appendChild(img);
            fileList.appendChild(listItem);
        })
        .catch(error => console.error(error));
});
