SAVEPOINT test;
    DO $dollar$
    BEGIN
        BEGIN
            PERFORM meme_factory._generate_big_meme('bad-luck-brian', 1000, 1000);
        EXCEPTION
        WHEN raise_exception THEN
            IF SQLERRM IS DISTINCT FROM 'unlucky' THEN
                RAISE EXCEPTION 'Expected exception "unlucky", got: % %', SQLSTATE, SQLERRM;
            END IF;
            RETURN;
        END;
        RAISE EXCEPTION 'Expect exception "unlucky" for generation of meme "bad-luck-brian".';
    END;
    $dollar$;
ROLLBACK TO SAVEPOINT test;
