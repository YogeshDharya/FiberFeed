// public/script.js
document.addEventListener('DOMContentLoaded', function () {
    axios.get('/news')
        .then(function (response) {
            const newsContainer = document.getElementById('news');
            response.data.news.forEach(article => {
                const newsItem = document.createElement('div');
                newsItem.className = 'news-item';

                const title = document.createElement('div');
                title.className = 'news-title';
                title.textContent = article.title;
                newsItem.appendChild(title);

                const description = document.createElement('div');
                description.className = 'news-description';
                description.textContent = article.description;
                newsItem.appendChild(description);

                const link = document.createElement('a');
                link.className = 'news-link';
                link.href = article.url;
                link.textContent = 'Read more';
                link.target = '_blank';
                newsItem.appendChild(link);

                newsContainer.appendChild(newsItem);
            });
        })
        .catch(function (error) {
            console.error('Error fetching news:', error);
        });
});
