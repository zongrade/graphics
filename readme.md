Генерация случайных данных по покупкам и просмотрам.\
Как это работает?\
Генерируется функция путём перемножения синусов с разным смещением. Получается вот так:\
![Случайная функция](https://github.com/zongrade/graph/raw/main/sin_product.png)\
затем берутся значения точек на этой функции. Эти значения это процентный рост от дельты предыдущего и минимума/максимума\
Чем больше синусов будет перемножено, тем более плоской будет функция\
Пример сгенерированных данных\
покупатели:\
![Случайные покупатели](https://github.com/zongrade/graph/raw/main/buyers_new.png)\
посетители:\
![Случайные посетители](https://github.com/zongrade/graph/raw/main/viewers.png)\
конверсия:\
![Случайная конверсия](https://github.com/zongrade/graph/raw/main/conversion.png)\
доход:\
![Случайный доход](https://github.com/zongrade/graph/raw/main/income.png)\
Все функции можно сгладить, если уменьшить количество значений по оси абсцисс