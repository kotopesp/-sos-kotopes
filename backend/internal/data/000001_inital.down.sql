DROP TABLE IF EXISTS "users" CASCADE;
DROP TABLE IF EXISTS "favourite_persons" CASCADE;
DROP TABLE IF EXISTS "favourite_posts" CASCADE;
DROP TABLE IF EXISTS "favourite_comments" CASCADE;
DROP TABLE IF EXISTS "roles_users" CASCADE;
DROP TABLE IF EXISTS "seekers" CASCADE;
DROP TABLE IF EXISTS "keepers" CASCADE;
DROP TABLE IF EXISTS "vets" CASCADE;
DROP TABLE IF EXISTS "vet_reviews" CASCADE;
DROP TABLE IF EXISTS "keeper_reviews" CASCADE;
DROP TYPE IF EXISTS "animal_type" CASCADE;
DROP TYPE IF EXISTS "animal_status" CASCADE;
DROP TYPE IF EXISTS "animal_gender" CASCADE;
DROP TABLE IF EXISTS "animals" CASCADE;
DROP TABLE IF EXISTS "messages" CASCADE;
DROP TABLE IF EXISTS "conversations" CASCADE;
DROP TABLE IF EXISTS "posts" CASCADE;
DROP TABLE IF EXISTS "post_response" CASCADE;
DROP TABLE IF EXISTS "post_likes" CASCADE;
DROP TABLE IF EXISTS "comments" CASCADE;
DROP TABLE IF EXISTS "comment_likes" CASCADE;