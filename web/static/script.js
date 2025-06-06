document.addEventListener('DOMContentLoaded', function () {
  document.querySelectorAll('.like').forEach(button => {
    button.addEventListener('click', function () {
      const vacancyElement = this.closest('.job');
      const vacancyId = vacancyElement.getAttribute('data-id');
      const isFavourite = this.getAttribute('data-fav') === 'true';

      const url = '/api/favourites';
      const method = isFavourite ? 'DELETE' : 'POST';

      fetch(url, {
        method: method,
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ vacancyId })
      })
      .then(response => {
        if (!response.ok) throw new Error('Ошибка при обновлении');
        return response.json();
      })
      .then(data => {
        if (isFavourite) {
          this.textContent = '♡';
          this.setAttribute('data-fav', 'false');
        } else {
          this.textContent = '♥';
          this.setAttribute('data-fav', 'true');
        }
      })
      .catch(error => {
        console.error('Ошибка:', error);
        alert('Не удалось обновить избранное');
      });
    });
  });
});
