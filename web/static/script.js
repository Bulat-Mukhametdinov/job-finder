document.addEventListener('DOMContentLoaded', function () {
  // Обработка лайков
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

  // Обработка комментариев
  document.querySelectorAll('.comment-section').forEach(section => {
    const vacancyId = section.getAttribute('data-vacancy-id');
    const commentDisplay = section.querySelector('.comment-display');
    const commentForm = section.querySelector('.comment-form');
    const commentText = section.querySelector('.comment-text');
    const commentInput = section.querySelector('.comment-input');
    
    const editBtn = section.querySelector('.edit-comment-btn');
    const saveBtn = section.querySelector('.save-comment-btn');
    const cancelBtn = section.querySelector('.cancel-comment-btn');

    // Кнопка "Edit" - переключение в режим редактирования
    if (editBtn) {
      editBtn.addEventListener('click', function() {
        commentDisplay.style.display = 'none';
        commentForm.style.display = 'block';
        commentInput.focus();
      });
    }

    // Кнопка "Cancel" - отмена редактирования
    if (cancelBtn) {
      cancelBtn.addEventListener('click', function() {
        // Восстанавливаем исходное значение
        commentInput.value = commentText.textContent;
        commentForm.style.display = 'none';
        commentDisplay.style.display = 'block';
      });
    }

    // Кнопка "Save" - сохранение комментария
    if (saveBtn) {
      saveBtn.addEventListener('click', function() {
        const comment = commentInput.value.trim();
        
        // Создаем URLSearchParams для отправки form data
        const formData = new URLSearchParams();
        formData.append('vacancyId', vacancyId);
        formData.append('comment', comment);

        fetch('/api/comment-favourite', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
          },
          body: formData
        })
        .then(response => {
          console.log('Response status:', response.status);
          if (!response.ok) {
            return response.text().then(text => {
              console.log('Error response:', text);
              throw new Error(`Ошибка при сохранении: ${response.status} - ${text}`);
            });
          }
          return response.json();
        })
        .then(data => {
          console.log('Success response:', data);
          if (data.success) {
            // Если комментарий пустой, оставляем в режиме создания
            if (comment === '') {
              commentInput.placeholder = 'Leave your comments...';
              return;
            }

            // Ищем или создаем элементы для отображения
            let commentDisplay = section.querySelector('.comment-display');
            let commentTextEl = section.querySelector('.comment-text');
            
            if (!commentDisplay) {
              // Создаем новый блок для отображения комментария
              commentDisplay = document.createElement('div');
              commentDisplay.className = 'comment-display';
              commentDisplay.innerHTML = `
                <p class="comment-text">${comment}</p>
                <button type="button" class="edit-comment-btn">Edit</button>
              `;
              section.insertBefore(commentDisplay, commentForm);
              
              // Добавляем кнопку Cancel в форму, если её нет
              if (!cancelBtn) {
                const newCancelBtn = document.createElement('button');
                newCancelBtn.type = 'button';
                newCancelBtn.className = 'cancel-comment-btn';
                newCancelBtn.textContent = 'Cancel';
                commentForm.appendChild(newCancelBtn);
                
                // Добавляем обработчик для новой кнопки Cancel
                newCancelBtn.addEventListener('click', function() {
                  const currentComment = commentDisplay.querySelector('.comment-text').textContent;
                  commentInput.value = currentComment;
                  commentForm.style.display = 'none';
                  commentDisplay.style.display = 'block';
                });
              }
              
              // Добавляем обработчик для кнопки Edit
              const newEditBtn = commentDisplay.querySelector('.edit-comment-btn');
              newEditBtn.addEventListener('click', function() {
                commentInput.value = commentDisplay.querySelector('.comment-text').textContent;
                commentDisplay.style.display = 'none';
                commentForm.style.display = 'block';
                commentInput.focus();
              });
            } else {
              // Обновляем существующий комментарий
              commentTextEl.textContent = comment;
            }

            // Переключаемся в режим просмотра
            commentForm.style.display = 'none';
            commentDisplay.style.display = 'block';
          }
        })
        .catch(error => {
          console.error('Ошибка:', error);
          alert('Не удалось сохранить комментарий');
        });
      });
    }

    // Обработка Enter в поле ввода
    if (commentInput) {
      commentInput.addEventListener('keypress', function(e) {
        if (e.key === 'Enter') {
          saveBtn.click();
        }
      });
    }
  });
});