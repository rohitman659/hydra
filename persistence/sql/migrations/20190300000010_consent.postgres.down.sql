DROP INDEX hydra_oauth2_logout_request_veri_idx;
ALTER TABLE hydra_oauth2_logout_request DROP CONSTRAINT hydra_oauth2_logout_request_client_id_fk;
