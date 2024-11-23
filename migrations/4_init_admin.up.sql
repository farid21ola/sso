INSERT INTO users(id, email, pass_hash, is_admin)
VALUES (1, 'admin@admin.com', '$2a$10$LkILTxGhAgjZ8DkJBAGIzOyTWe1GbluCJI4gGTcTeBlsActhYufvq', true)
    ON CONFLICT DO NOTHING;