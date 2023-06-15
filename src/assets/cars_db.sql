CREATE TABLE listed_cars (
    manufacturer TEXT,
    model TEXT,
    stock INT
);

CREATE TABLE wanted_cars (
    manufacturer TEXT,
    model TEXT
);

INSERT INTO listed_cars(manufacturer, model, stock) VALUES('Ford', 'Mustang GT', 3);
INSERT INTO listed_cars(manufacturer, model, stock) VALUES('Dodge', 'Charger Hellcat', 4);
INSERT INTO listed_cars(manufacturer, model, stock) VALUES('Chevrolet', 'El Camino', 3);

INSERT INTO wanted_cars(manufacturer, model) VALUES('Toyota', 'Supra Mk4');
INSERT INTO wanted_cars(manufacturer, model) VALUES('Honda', 'Civic Type R');
INSERT INTO wanted_cars(manufacturer, model) VALUES('Porsche', '911 Carrera');
