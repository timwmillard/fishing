
-- Murrabit Cod Challenge
INSERT INTO fishing.club(
    name,
    slug,
    legal_name
)
VALUES (
    'Murrabit Cod Challenge',
    'murrabitcodchallenge',
    'Murrabit Football & Netball Club Inc'
);

-- 2022
INSERT INTO fishing.event(slug,name,start_timestamp,end_timestamp,location,club_id)
VALUES
('2009','Murrabit Cod Challenge 2009',NULL,NULL,'Murrabit Recreation Reserve',1),
('2010','Murrabit Cod Challenge 2010',NULL,NULL,'Murrabit Recreation Reserve',1),
('2011','Murrabit Cod Challenge 2011',NULL,NULL,'Murrabit Recreation Reserve',1),
('2012','Murrabit Cod Challenge 2012',NULL,NULL,'Murrabit Recreation Reserve',1),
('2013','Murrabit Cod Challenge 2013',NULL,NULL,'Murrabit Recreation Reserve',1),
('2014','Murrabit Cod Challenge 2014',NULL,NULL,'Murrabit Recreation Reserve',1),
('2015','Murrabit Cod Challenge 2015',NULL,NULL,'Murrabit Recreation Reserve',1),
('2016','Murrabit Cod Challenge 2016',NULL,NULL,'Murrabit Recreation Reserve',1),
('2017','Murrabit Cod Challenge 2017',NULL,NULL,'Murrabit Recreation Reserve',1),
('2018','Murrabit Cod Challenge 2018',NULL,NULL,'Murrabit Recreation Reserve',1),
('2019','Murrabit Cod Challenge 2019',NULL,NULL,'Murrabit Recreation Reserve',1),
('2021','Murrabit Cod Challenge 2021',NULL,NULL,'Murrabit Recreation Reserve',1),
('2022','Murrabit Cod Challenge 2022',NULL,NULL,'Murrabit Recreation Reserve',1),
('2023','Murrabit Cod Challenge 2023',NULL,NULL,'Murrabit Recreation Reserve',1);

UPDATE fishing.club
SET current_event = 14
WHERE id = 1;

-- Species
INSERT INTO fishing.species(common_name,scientific_name,slug,photo_url)
VALUES
('Murray Cod','Maccullochella peelii','murraycod',NULL),
('Yellow Belly','Macquaria ambigua','yellowbelly',NULL),
('Carp','Cyprinus carpio','carp',NULL),
('Trout Cod','Maccullochella macquariensis','troutcod',NULL);
