INSERT INTO memberships (id, name, description, scopes) VALUES
    ('3463d54a-68fa-4877-aa5b-2d05f39613bb', 'Free', 'Perfect for free trial', '*:currencies *:items *:warehouses'),
    ('df7b5f6b-7914-49da-8fc8-d4a4fc710461', 'Premium', 'Perfect for paid users', '*:currencies *:kardex *:items *:warehouses'),
    ('5faf08c0-69f1-40d4-bb3d-a0938f797d66', 'Admin', 'Internal use only', '*:*');

-- primary company will be premium tier, secondary free tier.
INSERT INTO company_memberships (company_id, membership_id) VALUES
    ('b918deaf-92ab-485d-9a69-ee7a2a5f4aef', 'df7b5f6b-7914-49da-8fc8-d4a4fc710461'),
    ('cde98fd2-6023-4953-9e95-884aed1f09ce', '3463d54a-68fa-4877-aa5b-2d05f39613bb');