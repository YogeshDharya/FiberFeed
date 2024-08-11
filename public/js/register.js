document.addEventListener('DOMContentLoaded', function () {
    const token = localStorage.getItem('authToken');
    if (!token) {
        window.location.href = '/login.html';  // Redirect to login if no token
        return;
    }

    axios.get('/news', {
        headers: {
            'Authorization': `Bearer ${token}`  // Attach JWT token to the request
        }
    })
    .then(function (response) {
        const newsContainer = document.getElementById('news');
        response.data.news.forEach(article => {
            const newsItem = document.createElement('div');
            newsItem.className = 'news-item';

            const title = document.createElement('h2');
            title.textContent = article.title;
            newsItem.appendChild(title);

            const description = document.createElement('p');
            description.textContent = article.description;
            newsItem.appendChild(description);

            if (article.image) {
                const image = document.createElement('img');
                image.className = 'news-image';
                image.src = article.image;
                image.alt = article.title;
                newsItem.appendChild(image);
            }

            const author = document.createElement('p');
            author.textContent = `Author: ${article.author}`;
            newsItem.appendChild(author);

            const readMore = document.createElement('p');
            const link = document.createElement('a');
            link.href = article.url;
            link.textContent = 'Read More';
            link.target = '_blank';
            readMore.appendChild(link);
            newsItem.appendChild(readMore);

            newsContainer.appendChild(newsItem);
        });
    })
    .catch(function (error) {
        console.error('Error fetching news:', error);
        alert('Failed to fetch news. Please try again later.');
    });
});
