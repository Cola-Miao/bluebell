package rdb

import (
	"bluebell/model"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func CreateArticle(uuid int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := db.ZAdd(ctx, formatKey(articleCreateAt), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: uuid,
	}).Err(); err != nil {
		return err
	}
	if err := db.ZAdd(ctx, formatKey(articleScore), redis.Z{
		Score:  0,
		Member: uuid,
	}).Err(); err != nil {
		return err
	}
	return nil
}

func HotArticle(uuid string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := db.ZScore(ctx, formatKey(articleCreateAt), uuid).Result()
	if err != nil {
		return false, fmt.Errorf("%w: %w", model.ErrGetCache, err)
	}
	t := time.Unix(int64(res), 0).Add(articleExpire)
	if t.Before(time.Now()) {
		return false, nil
	}
	return true, nil
}

func VoteForArticle(artID, userID string, score float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	pp := db.TxPipeline()

	oldScore, err := db.ZScore(ctx, formatKey(articleVoter, artID), userID).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("%w: %w", model.ErrGetCache, err)
	}

	if err = pp.ZAdd(ctx, formatKey(articleVoter, artID), redis.Z{
		Score:  score,
		Member: userID,
	}).Err(); err != nil {
		return fmt.Errorf("%w: %w", model.ErrSetCache, err)
	}
	if err = pp.ZIncrBy(ctx, formatKey(articleScore), score-oldScore, artID).Err(); err != nil {
		return fmt.Errorf("%w: %w", model.ErrSetCache, err)
	}
	if _, err = pp.Exec(ctx); err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("transaction failed: %w", err)
	}
	return nil
}

func ArticleScore(uuid string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	score, err := db.ZScore(ctx, formatKey(articleScore), uuid).Result()
	if err != nil {
		return 0, fmt.Errorf("%w: %w", model.ErrGetCache, err)
	}
	voters, err := db.ZCard(ctx, formatKey(articleVoter, uuid)).Result()
	if err != nil || voters == 0 {
		return 0, fmt.Errorf("%w: %w", model.ErrGetCache, err)
	}
	return score / float64(voters), nil
}

func DeleteArticle(uuid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var err error
	//if err = db.ZRem(ctx, formatKey(articleScore), uuid).Err(); err != nil {
	//	return fmt.Errorf("delete cache item failed: %w", err)
	//}
	if err = db.ZRem(ctx, formatKey(articleCreateAt), uuid).Err(); err != nil {
		return fmt.Errorf("delete cache item failed: %w", err)
	}
	if err = db.Del(ctx, formatKey(articleVoter, uuid)).Err(); err != nil {
		return fmt.Errorf("delete cache failed: %w", err)
	}
	return nil
}

func HighestScoreArticle(offset, size int64) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := db.ZRevRangeByScore(ctx, formatKey(articleScore), &redis.ZRangeBy{
		Offset: offset,
		Count:  size,
	}).Result()
	return res, err
}
