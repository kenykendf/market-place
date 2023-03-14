-- INSERT QUERY CATEGORIES
INSERT INTO public.categories
(id, created_at, updated_at, deleted_at, categories_name)
VALUES(1, '2023-03-14 02:40:18.130', '2023-03-14 02:40:18.130', NULL, 'kitchenware');
INSERT INTO public.categories
(id, created_at, updated_at, deleted_at, categories_name)
VALUES(2, '2023-03-14 14:32:34.520', '2023-03-14 14:32:34.520', NULL, 'beverages');
INSERT INTO public.categories
(id, created_at, updated_at, deleted_at, categories_name)
VALUES(3, '2023-03-14 14:37:46.236', '2023-03-14 14:37:46.236', NULL, 'foods');

-- INSERT QUERY ITEMS
INSERT INTO public.items
(id, created_at, updated_at, deleted_at, item_code, item_name, description, quantity, categories_id)
VALUES(1, '2023-03-14 14:40:25.823', '2023-03-14 14:40:25.823', NULL, 'BVG001', 'Coca Cola', 'Soft Drink', 100, 2);
INSERT INTO public.items
(id, created_at, updated_at, deleted_at, item_code, item_name, description, quantity, categories_id)
VALUES(2, '2023-03-14 14:40:29.835', '2023-03-14 14:40:29.835', NULL, 'BVG002', 'Fanta', 'Soft Drink', 100, 2);
INSERT INTO public.items
(id, created_at, updated_at, deleted_at, item_code, item_name, description, quantity, categories_id)
VALUES(3, '2023-03-14 14:40:32.334', '2023-03-14 14:40:32.334', NULL, 'BVG003', 'Sprite', 'Soft Drink', 100, 2);
INSERT INTO public.items
(id, created_at, updated_at, deleted_at, item_code, item_name, description, quantity, categories_id)
VALUES(4, '2023-03-14 14:42:20.571', '2023-03-14 14:42:20.571', NULL, 'KCW001', 'Pan', 'Frying Pan', 5, 1);
INSERT INTO public.items
(id, created_at, updated_at, deleted_at, item_code, item_name, description, quantity, categories_id)
VALUES(5, '2023-03-14 14:42:47.698', '2023-03-14 14:42:47.698', NULL, 'KCW002', 'Spatula', 'Spatula', 5, 1);
INSERT INTO public.items
(id, created_at, updated_at, deleted_at, item_code, item_name, description, quantity, categories_id)
VALUES(6, '2023-03-14 14:43:45.055', '2023-03-14 14:43:45.055', NULL, 'FOO001', 'Fiesta Chicken Nuggets', 'Frozen Foods', 5, 3);
INSERT INTO public.items
(id, created_at, updated_at, deleted_at, item_code, item_name, description, quantity, categories_id)
VALUES(7, '2023-03-14 14:44:22.755', '2023-03-14 14:44:22.755', NULL, 'FOO002', 'Sosis So Nice', 'Sausage', 50, 3);
INSERT INTO public.items
(id, created_at, updated_at, deleted_at, item_code, item_name, description, quantity, categories_id)
VALUES(8, '2023-03-14 14:45:17.157', '2023-03-14 14:45:17.157', NULL, 'FOO003', 'Breast Chicken Fillet', 'Raw Chicken', 10, 3);


