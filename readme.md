Programa que consume una API de Finnhub, para obtener en tiempo real el precio actual, volumen, precio de cierre, entre muchas otras variables de una criptomoneda, acción, par de divisas, etc.  

Además, le agregué un indicador en base a las medias móviles simples de 10 y 20.  El resultado de una media móvil simple es el precio de cierre de los últimos 10 y 20 días, divido esa misma cantidad de días. 

El programa manda una señal de compra cuando la media móvil rápida (de 10 días) supere a la lenta (de 20 días), y una señal de venta en caso de que suceda lo contrario. Ejecuta ordenes de compra y venta en binance consumiendo la API de dicha plataforma. 

Tiene un conversor de precios (de usdt a btc, por ejemplo). Permite setear stop loss y take profit. El programa envía todas las ordenes ejecutadas a una base de datos en MySQL, utilizando el framework GORM.

Este programa va a ser utilizado para un algoritmo de Machine Learning real que enviará alertas reales y mucho más certeras.