-- +goose Up
-- +goose StatementBegin

ALTER TABLE production_companies
    DROP CONSTRAINT production_companies_origin_country_fkey, 
    ADD CONSTRAINT production_companies_origin_country_fkey FOREIGN KEY (origin_country) REFERENCES countries(iso_3166_1); 

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE production_companies
    DROP CONSTRAINT production_companies_origin_country_fkey,  
    ADD CONSTRAINT production_companies_origin_country_fkey FOREIGN KEY (origin_country) REFERENCES origin_countries(iso_3166_1);  

-- +goose StatementEnd