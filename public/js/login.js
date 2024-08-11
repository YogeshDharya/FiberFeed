document.getElementById('loginButton').addEventListener('click', function () {
    authenticateUser('/login');
});

document.getElementById('registerButton').addEventListener('click', function () {
    authenticateUser('/register');
});

function authenticateUser(url) {
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    axios.post(url, {
        username: username,
        password: password
    })
    .then(function (response) {
        const token = response.data.token;
        localStorage.setItem('authToken', token);  // Store JWT token in local storage
        window.location.href = '/news.html';  // Redirect to news aggregator homepage
    })
    .catch(function (error) {
        console.error('Authentication failed:', error);
        alert('Authentication failed. Please check your credentials.');
    });
}
