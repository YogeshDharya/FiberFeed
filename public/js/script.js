document.addEventListener('DOMContentLoaded', function () {
    axios.get('/news')
        .then(function (response) {
            const newsContainer = document.getElementById('news');
            response.data.news.forEach(article => {
                const newsItem = document.createElement('div');
                newsItem.className = 'news-item';

                const title = document.createElement('h2');
                title.className = 'news-title';
                title.textContent = article.title;
                newsItem.appendChild(title);

                const description = document.createElement('p');
                description.className = 'news-description';
                description.textContent = article.description;
                newsItem.appendChild(description);

                const link = document.createElement('a');
                link.className = 'news-link';
                link.textContent = 'Read More';
                link.href = article.url;
                link.target = '_blank';
                newsItem.appendChild(link);

                const author = document.createElement('p');
                author.className = 'news-author';
                author.textContent = article.author;
                newsItem.appendChild(author);

                const image = document.createElement('img');
                image.className = 'news-image';
                image.src = article.image; 
                image.alt = article.title;
                newsItem.appendChild(image);

                const newsId = document.createElement('span');
                newsId.className = 'news-id';
                newsId.textContent = `ID: ${article.id}`;
                newsItem.appendChild(newsId);

                newsContainer.appendChild(newsItem);
            });
        })
        .catch(function (error) {
            console.error('Error fetching news:', error);
        });
});
