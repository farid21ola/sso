INSERT INTO apps(id, name, secret)
VALUES (1, 'test', 'top_secret')
ON CONFLICT DO NOTHING;