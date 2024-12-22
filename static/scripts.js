async function addBook() {
    const title = document.getElementById("title").value;
    const author = document.getElementById("author").value;
    const year = parseInt(document.getElementById("year").value);

    const response = await fetch("/api/books", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title, author, year })
    });

    const result = await response.json();
    alert(result.message);
}

async function listBooks() {
    const response = await fetch("/api/books");
    const books = await response.json();
    const list = document.getElementById("book-list");
    list.innerHTML = "";
    books.forEach(book => {
        const item = document.createElement("li");
        item.textContent = `${book.title}, por ${book.author} (${book.year})`;
        list.appendChild(item);
    });
}

async function deleteBook() {
    const title = document.getElementById("delete-title").value;
    const response = await fetch(`/api/books?title=${title}`, {
        method: "DELETE"
    });

    const result = await response.json();
    alert(result.message);
}
