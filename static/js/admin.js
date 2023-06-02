document.querySelectorAll("input").forEach(function (input) {
    input.addEventListener("change", () => {
        updateCardPreview();
    });
});

function updateCardPreview() {
    if (document.getElementById("post-title").value)
        document.getElementById("card-title").textContent = document.getElementById("post-title").value;
    if (document.getElementById("post-description").value)
        document.getElementById("card-subtitle").textContent = document.getElementById("post-description").value;
    if (document.getElementById("post-author").value)
        document.getElementById("card-author").textContent = document.getElementById("post-author").value;
    if (document.getElementById("post-date").value)
        document.getElementById("card-date").textContent = document.getElementById("post-date").value;
    if (document.getElementById("post-title").value)
        document.getElementById("article-title").textContent = document.getElementById("post-title").value;
    if (document.getElementById("post-description").value)
        document.getElementById("article-subtitle").textContent = document.getElementById("post-description").value;
}

function uploadAvatar() {
    const avatar_form = document.getElementById("author-avatar-form");
    const author_avatar = document.getElementById("author-avatar");
    const card_avatar = document.getElementById("card-avatar");
    const file = document.getElementById("author-photo-img").files[0];
    const reader = new FileReader();

    reader.addEventListener("load", () => {
        author_avatar.src = reader.result;
        card_avatar.src = reader.result;
    });
    
    if (file) {
        reader.readAsDataURL(file);
    }

    avatar_form.querySelector(".form__upload").classList.add("hidden");
    avatar_form.querySelector(".form__photo-upload-new").classList.remove("hidden");
    avatar_form.querySelector(".form__photo-remove").classList.remove("hidden");
}

function removeAvatar(event) {
    event.preventDefault();

    const avatar_form = document.getElementById("author-avatar-form");
    const author_avatar = document.getElementById("author-avatar");
    const card_avatar = document.getElementById("card-avatar");

    author_avatar.src = "/static/img/camera.png";
    card_avatar.src = "/static/img/empty.png";

    avatar_form.querySelector(".form__upload").classList.remove("hidden");
    avatar_form.querySelector(".form__photo-upload-new").classList.add("hidden");
    avatar_form.querySelector(".form__photo-remove").classList.add("hidden");
}

function uploadHeroBig() {
    const hero_form = document.getElementById("hero-image-big");
    const hero_image = hero_form.querySelector(".hero-image__picture_big");
    const article_image = document.getElementById("article-image");
    const file = document.getElementById("hero-picture-big").files[0];
    const reader = new FileReader();

    reader.addEventListener("load", () => {
        hero_image.src = reader.result;
        article_image.src = reader.result;
    });
    
    if (file) {
        reader.readAsDataURL(file);
    }

    hero_form.querySelector(".form__photo-upload-new").classList.remove("hidden");
    hero_form.querySelector(".form__photo-remove").classList.remove("hidden");
    hero_form.querySelector(".img-field-description").classList.add("hidden");
}

function removeHeroBig(event) {
    event.preventDefault();

    const hero_form = document.getElementById("hero-image-big");
    const hero_image = hero_form.querySelector(".hero-image__picture_big");
    const article_image = document.getElementById("article-image");

    hero_image.src = "/static/img/upload-hero-big.png";
    article_image.src = "/static/img/empty.png";

    hero_form.querySelector(".form__photo-upload-new").classList.add("hidden");
    hero_form.querySelector(".form__photo-remove").classList.add("hidden");
    hero_form.querySelector(".img-field-description").classList.remove("hidden");
}

function uploadHeroSmall() {
    const hero_form = document.getElementById("hero-image-small");
    const hero_image = hero_form.querySelector(".hero-image__picture_small");
    const card_image = document.getElementById("card-image");
    const file = document.getElementById("hero-picture-small").files[0];
    const reader = new FileReader();

    reader.addEventListener("load", () => {
        hero_image.src = reader.result;
        card_image.src = reader.result;
    });
    
    if (file) {
        reader.readAsDataURL(file);
    }

    hero_form.querySelector(".form__photo-upload-new").classList.remove("hidden");
    hero_form.querySelector(".form__photo-remove").classList.remove("hidden");
    hero_form.querySelector(".img-field-description").classList.add("hidden");
}

function removeHeroSmall(event) {
    event.preventDefault();

    const hero_form = document.getElementById("hero-image-small");
    const hero_image = hero_form.querySelector(".hero-image__picture_small");
    const card_image = document.getElementById("card-image");

    hero_image.src = "/static/img/upload-hero-small.png";
    card_image.src = "/static/img/empty.png";

    hero_form.querySelector(".form__photo-upload-new").classList.add("hidden");
    hero_form.querySelector(".form__photo-remove").classList.add("hidden");
    hero_form.querySelector(".img-field-description").classList.remove("hidden");
}

function getBase64(file) {
    return new Promise((resolve, reject) => {
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = () => resolve(reader.result.split(",")[1]);
      reader.onerror = (error) => reject(error);
    });
  }

async function publish(event) {
    event.preventDefault();

    const main_form = document.getElementsByClassName("main__form")[0];
    const formData = new FormData(main_form);
    const outData = {};

    for (let [key, value] of formData.entries()) {
      if (key == "author-photo-img" || key == "hero-picture-big" || key == "hero-picture-small") {
        outData[key] = await getBase64(value);
        outData[`${key}-file-name`] = value.name;
      } else {
        outData[key] = value;
      }
    }

    console.log(outData);

    const response = await fetch("/api/post", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    });

    console.log(response);
  }